package common

import (
	"github.com/SeeDAO-OpenSource/sgn/pkg/app"
	"github.com/SeeDAO-OpenSource/sgn/pkg/blob"
	"github.com/SeeDAO-OpenSource/sgn/pkg/blob/file"
	"github.com/SeeDAO-OpenSource/sgn/pkg/di"
	"github.com/SeeDAO-OpenSource/sgn/pkg/utils"
)

func AddFileSystemBlobStore(builder *app.AppBuilder) {
	builder.ConfigureServices(func() error {
		var fileOptions = &file.FileBlobStoreOptions{
			BasePath: "/data",
		}
		utils.ViperBind("FileBlobStore", fileOptions)
		di.AddValue(fileOptions)
		di.AddTransient(func(c *di.Container) *blob.BlobStore {
			blobStore := file.NewFileBlobStore(fileOptions)
			return &blobStore
		})
		return nil
	})
}
