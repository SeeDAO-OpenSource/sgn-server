package common

import (
	"github.com/SeeDAO-OpenSource/sgn/pkg/app"
	"github.com/SeeDAO-OpenSource/sgn/pkg/blob"
	"github.com/SeeDAO-OpenSource/sgn/pkg/blob/file"
	"github.com/SeeDAO-OpenSource/sgn/pkg/services"
)

func AddFileSystemBlobStore(builder *app.AppBuilder) {
	var fileOptions = &file.FileBlobStoreOptions{
		BasePath: "/data",
	}
	builder.BindOptions("FileBlobStore", fileOptions)
	builder.ConfigureServices(func() error {
		services.AddValue(fileOptions)
		services.AddTransient(func(c *services.Container) *blob.BlobStore {
			blobStore := file.NewFileBlobStore(fileOptions)
			return &blobStore
		})
		return nil
	})
}
