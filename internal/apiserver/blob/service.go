package blob

import "github.com/SeeDAO-OpenSource/sgn/pkg/blob"

type BlobService struct {
	store blob.BlobStore
}

func NewBlobService(store blob.BlobStore) *BlobService {
	return &BlobService{
		store: store,
	}
}

func (srv *BlobService) Get(key string, process *blob.Process) (*blob.BlobReader, error) {
	return srv.store.Read(&key, process)
}
