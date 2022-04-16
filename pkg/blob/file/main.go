package file

import (
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"

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
		os.Mkdir(f.options.BasePath, os.ModePerm)
		os.Chmod(f.options.BasePath, os.ModePerm)
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

func (f *FileBlobStore) Read(key *string) (*blob.BlobReader, error) {
	ukey := convertKey(*key)
	path := filepath.Join(f.options.BasePath, ukey)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadFile(path)
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
