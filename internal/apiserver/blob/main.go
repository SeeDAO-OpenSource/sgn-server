package blob

import (
	"github.com/SeeDAO-OpenSource/sgn/pkg/blob/file"
	"github.com/SeeDAO-OpenSource/sgn/pkg/di"
	"github.com/SeeDAO-OpenSource/sgn/pkg/server"
)

func BlobStore(buider *server.ServerBuiler) error {
	buider.Configure(initRoute)
	buider.App.ConfigureServices(func() error {
		di.AddTransient(func(c *di.Container) *BlobService {
			fileOptions := di.Get[file.FileBlobStoreOptions]()
			blobStore := file.NewFileBlobStore(fileOptions)
			blobService := NewBlobService(blobStore)
			return blobService
		})
		return nil
	})
	return nil
}

func initRoute(s *server.Server) error {
	sgnCtl := newBlobController()
	route(&sgnCtl, s.G)
	return nil
}
