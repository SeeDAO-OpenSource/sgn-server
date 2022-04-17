package file

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path/filepath"

	"github.com/nfnt/resize"
	"github.com/waite-lee/nftserver/pkg/blob"
)

type FileBlobStoreOptions struct {
	BasePath string
}

type FileBlobStore struct {
	options *FileBlobStoreOptions
}

func NewFileBlobStore(options *FileBlobStoreOptions) blob.BlobStore {
	return &FileBlobStore{
		options: options,
	}
}

func (f *FileBlobStore) Save(key *string, content *[]byte, overwrite bool) error {
	ukey := convertKey(*key)
	_, err := os.Stat(f.options.BasePath)
	if os.IsNotExist(err) {
		dir, err := filepath.Abs(f.options.BasePath)
		if err != nil {
			return err
		}
		os.Mkdir(dir, os.ModePerm)
	}
	if overwrite || !isExsits(filepath.Join(f.options.BasePath, ukey)) {
		path := filepath.Join(f.options.BasePath, ukey)
		err = ioutil.WriteFile(path, *content, 0644)
	}
	return err
}
func (f *FileBlobStore) Exists(key *string) bool {
	ukey := convertKey(*key)
	path := filepath.Join(f.options.BasePath, ukey)
	return isExsits(path)
}

func (f *FileBlobStore) Read(key *string, process *blob.Process) (*blob.BlobReader, error) {
	ukey := convertKey(*key)
	path := filepath.Join(f.options.BasePath, ukey)
	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	content, err := processFile(path, process)
	if err != nil {
		return nil, err
	}
	return &blob.BlobReader{
		Content: content,
		Size:    fileInfo.Size(),
		Name:    fileInfo.Name(),
	}, nil
}

func isExsits(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func convertKey(key string) string {
	return url.QueryEscape(key)
}

func processFile(path string, process *blob.Process) ([]byte, error) {
	var content []byte
	var err error
	if process == nil {
		content, err = ioutil.ReadFile(path)
	}
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	img, format, err := image.Decode(file)
	log.Println(format)
	if err == nil {
		if process.Width > 0 && process.Height > 0 {
			m := resize.Resize(uint(process.Width), uint(process.Height), img, resize.Lanczos3)
			var buf bytes.Buffer
			switch format {
			case "png":
				err = png.Encode(&buf, m)
			case "jpeg":
				err = jpeg.Encode(&buf, m, nil)
			}
			content = buf.Bytes()
		}
	}
	if len(content) == 0 {
		content, err = ioutil.ReadFile(path)
	}
	return content, err
}
