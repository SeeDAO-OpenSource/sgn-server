package common

import "github.com/SeeDAO-OpenSource/sgn/pkg/app"

func CommonServices(builder *app.AppBuilder) {
	builder.Use(AddIpfsClient).
		Use(AddEthClient).
		Use(AddFileSystemBlobStore).
		Use(AddHttpClient).
		Use(AddMongoClient)
}
