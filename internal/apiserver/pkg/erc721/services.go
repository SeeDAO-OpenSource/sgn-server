package erc721

import (
	"github.com/SeeDAO-OpenSource/sgn/pkg/app"
	"github.com/SeeDAO-OpenSource/sgn/pkg/mvc"
	"github.com/SeeDAO-OpenSource/sgn/pkg/services"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nanmu42/etherscan-api"
)

func AddErc721Services(builder *app.AppBuilder) {
	var esOptions = &EtherScanOptions{
		BaseURL: "https://api.etherscan.io/api?",
	}
	builder.BindOptions("EtherScan", esOptions)
	builder.ConfigureServices(func() error {
		services.AddValue(esOptions)
		services.AddTransient(func(c *services.Container) *etherscan.Client {
			options := services.Get[EtherScanOptions]()
			httpOptions := services.Get[mvc.HttpClientOptions]()
			client, err := GetClient(options, httpOptions)
			if err != nil {
				return nil
			}
			return client
		})
		services.AddTransient(func(c *services.Container) *Erc721Service {
			client := services.Get[etherscan.Client]()
			options := services.Get[EtherScanOptions]()
			ethClient := services.Get[ethclient.Client]()
			return NewErc721Service(client, options, ethClient)
		})
		return nil
	})
}
