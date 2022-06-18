package blob

import (
	"github.com/SeeDAO-OpenSource/sgn/pkg/blob/file"
	"github.com/SeeDAO-OpenSource/sgn/pkg/server"
	"github.com/SeeDAO-OpenSource/sgn/pkg/services"
)

func AddBlob(buider *server.ServerBuiler) error {
	buider.Configure(initRoute)
	buider.App.ConfigureServices(func() error {
		services.AddTransient(func(c *services.Container) *BlobService {
			fileOptions := services.Get[file.FileBlobStoreOptions]()
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
