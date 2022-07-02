package erc721

import (
	"github.com/SeeDAO-OpenSource/sgn/pkg/app"
	"github.com/SeeDAO-OpenSource/sgn/pkg/di"
	"github.com/SeeDAO-OpenSource/sgn/pkg/mvc"
	"github.com/SeeDAO-OpenSource/sgn/pkg/utils"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nanmu42/etherscan-api"
)

func AddErc721Services(builder *app.AppBuilder) {

	builder.ConfigureServices(func() error {
		var esOptions = &EtherScanOptions{
			BaseURL: "https://api.etherscan.io/api?",
		}
		utils.ViperBind("EtherScan", esOptions)
		di.AddValue(esOptions)
		di.AddTransient(func(c *di.Container) *etherscan.Client {
			options := di.Get[EtherScanOptions]()
			httpOptions := di.Get[mvc.HttpClientOptions]()
			client, err := GetClient(options, httpOptions)
			if err != nil {
				return nil
			}
			return client
		})
		di.AddTransient(func(c *di.Container) *Erc721Service {
			client := di.Get[etherscan.Client]()
			options := di.Get[EtherScanOptions]()
			ethClient := di.Get[ethclient.Client]()
			return NewErc721Service(client, options, ethClient)
		})
		return nil
	})
}
