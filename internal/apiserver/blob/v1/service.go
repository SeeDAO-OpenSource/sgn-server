package blobv1

import "github.com/waite-lee/nftserver/pkg/blob"

type BlobService struct {
	store blob.BlobStore
}

func NewBlobService(store blob.BlobStore) *BlobService {
	return &BlobService{
		store: store,
	}
}

func (srv *BlobService) Get(key string) (*blob.BlobReader, error) {
	return srv.store.Read(&key)
}
