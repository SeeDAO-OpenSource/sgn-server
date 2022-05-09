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
	"strconv"
	"strings"

	"github.com/nfnt/resize"
	"github.com/waite-lee/sgn/pkg/blob"
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
	ukey := f.convertKey(*key)
	_, err := os.Stat(f.options.BasePath)
	if os.IsNotExist(err) {
		dir, err := filepath.Abs(f.options.BasePath)
		if err != nil {
			return err
		}
		os.MkdirAll(dir, os.ModePerm)
	}
	if overwrite || !f.isExsits(filepath.Join(f.options.BasePath, ukey)) {
		path := filepath.Join(f.options.BasePath, ukey)
		err = ioutil.WriteFile(path, *content, 0644)
	}
	return err
}

func (f *FileBlobStore) Exists(key *string) bool {
	ukey := f.convertKey(*key)
	path := filepath.Join(f.options.BasePath, ukey)
	return f.isExsits(path)
}

func (f *FileBlobStore) Read(key *string, process *blob.Process) (*blob.BlobReader, error) {
	ukey := f.convertKey(*key)
	path := filepath.Join(f.options.BasePath, ukey)
	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, nil
	}
	content, err := f.processFile(path, process)
	if err != nil {
		return nil, err
	}
	return &blob.BlobReader{
		Content: content,
		Size:    fileInfo.Size(),
		Name:    fileInfo.Name(),
	}, nil
}

func (f *FileBlobStore) isExsits(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func (f *FileBlobStore) convertKey(key string) string {
	return url.QueryEscape(key)
}

func (f *FileBlobStore) processFile(path string, process *blob.Process) ([]byte, error) {
	if process == nil {
		return ioutil.ReadFile(path)
	}
	cachPath := f.getCachPath(path, process)
	if cachPath != "" {
		content, err := os.ReadFile(cachPath)
		if err == nil {
			return content, nil
		}
	}
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	content, err := f.resizeImage(file, process)
	if err != nil {
		content, err = ioutil.ReadFile(path)
	}
	err = os.WriteFile(cachPath, content, 0644)
	return content, err
}

func (f *FileBlobStore) resizeImage(file *os.File, process *blob.Process) ([]byte, error) {
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
			return buf.Bytes(), nil
		}
	}
	return nil, err
}

func (f *FileBlobStore) getCachPath(path string, process *blob.Process) string {
	if process.Width > 0 && process.Height > 0 {
		ext := filepath.Ext(path)
		fileName := strings.TrimSuffix(filepath.Base(path), ext)
		dir := filepath.Join(filepath.Dir(path), ".cache", fileName)
		err := os.MkdirAll(dir, os.ModePerm)
		if err == nil {
			return filepath.Join(dir, strconv.Itoa(process.Width)+"_"+strconv.Itoa(process.Height)+ext)
		}
	}
	return ""
}
