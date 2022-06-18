package common

import (
	"github.com/SeeDAO-OpenSource/sgn/pkg/app"
	"github.com/SeeDAO-OpenSource/sgn/pkg/eth"
	"github.com/SeeDAO-OpenSource/sgn/pkg/services"
	"github.com/ethereum/go-ethereum/ethclient"
)

func AddEthClient(builder *app.AppBuilder) {
	var ethOptions = eth.NewEthOptions()
	builder.BindOptions("EthClient", ethOptions)
	builder.ConfigureServices(func() error {
		services.AddValue(ethOptions)
		services.AddTransient(func(c *services.Container) *ethclient.Client {
			options := services.Get[eth.EthOptions]()
			client, err := eth.GetClient(options)
			if err != nil {
				return nil
			}
			return client
		})
		return nil
	})
}
