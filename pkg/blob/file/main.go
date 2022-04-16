package file

import (
	"io/ioutil"
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
	_, err := os.Stat(f.options.BasePath)
	if os.IsNotExist(err) {
		os.Mkdir(f.options.BasePath, os.ModePerm)
		os.Chmod(f.options.BasePath, os.ModePerm)
	}
	if overwrite || !isExsits(filepath.Join(f.options.BasePath, *key)) {
		path := filepath.Join(f.options.BasePath, *key)
		err = ioutil.WriteFile(path, *content, 0644)
	}
	return err
}
func (f *FileBlobStore) Exists(key *string) bool {
	path := filepath.Join(f.options.BasePath, *key)
	return isExsits(path)
}

func isExsits(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
