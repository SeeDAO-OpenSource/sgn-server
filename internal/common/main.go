package common

import (
	"github.com/google/wire"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/waite-lee/sgn/pkg/app"
	"github.com/waite-lee/sgn/pkg/blob/file"
	"github.com/waite-lee/sgn/pkg/db/mongodb"
	"github.com/waite-lee/sgn/pkg/eth"
	"github.com/waite-lee/sgn/pkg/ipfs"
	"github.com/waite-lee/sgn/pkg/mvc"
)

var CommonSet = wire.NewSet(
	ipfs.GetClient,
	file.NewFileBlobStore,
	mvc.NewRequestClient,
	mongodb.GetClient,
	wire.Value(IpfsOptions),
	wire.Value(HttpOptions),
	wire.Value(MongoOptions),
	wire.Value(FileOptions),
	wire.Value(EthOptions),
)
var IpfsOptions = &ipfs.IpfsOptions{
	BaseURL: "https://ipfs.io/ipfs/",
}

var HttpOptions = &mvc.HttpClientOptions{}

var MongoOptions = &mongodb.MongoOptions{
	URL: "mongodb://localhost:27017",
}

var FileOptions = &file.FileBlobStoreOptions{
	BasePath: "data",
}

var EthOptions = eth.NewEthOptions()

func AddCommonOptions(ac *app.AppContext) {
	ac.CmdBuilder.PreRun(func(cmd *cobra.Command) error {
		initOptions()
		return nil
	})
}

func initOptions() {
	viper.UnmarshalKey("Ipfs", IpfsOptions)
	viper.UnmarshalKey("HttpClient", HttpOptions)
	viper.UnmarshalKey("Mongo", MongoOptions)
	viper.UnmarshalKey("BlobStore", FileOptions)
	viper.UnmarshalKey("EthClient", EthOptions)
}
