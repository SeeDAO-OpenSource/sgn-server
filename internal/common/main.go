package common

import (
	"github.com/google/wire"
	"github.com/waite-lee/nftserver/pkg/blob/file"
	"github.com/waite-lee/nftserver/pkg/db/mongodb"
	"github.com/waite-lee/nftserver/pkg/ipfs"
	"github.com/waite-lee/nftserver/pkg/mvc"
)

var IpfsOptions = &ipfs.IpfsOptions{
	BaseURL: "https://ipfs.io/ipfs/",
}

var HttpOptions = &mvc.HttpClientOptions{
	ProxyURL: "http://localhost:4780",
}

var MongoOptions = &mongodb.MongoOptions{
	URL: "mongodb://localhost:27017",
}

var FileOptions = &file.FileBlobStoreOptions{
	BasePath: "D:/data",
}

var CommonSet = wire.NewSet(
	ipfs.GetClient,
	file.NewFileBlobStore,
	mvc.NewRequestClient,
	mongodb.GetClient,
	wire.Value(IpfsOptions),
	wire.Value(HttpOptions),
	wire.Value(MongoOptions),
	wire.Value(FileOptions),
)
