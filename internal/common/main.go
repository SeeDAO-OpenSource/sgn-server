package common

import "github.com/SeeDAO-OpenSource/sgn/pkg/app"

func AddCommonServices(builder *app.AppBuilder) {
	AddIpfsClient(builder)
	AddEthClient(builder)
	AddFileSystemBlobStore(builder)
	AddHttpClient(builder)
	AddMongoClient(builder)
}
