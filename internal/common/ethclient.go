package common

import (
	"github.com/SeeDAO-OpenSource/sgn/pkg/app"
	"github.com/SeeDAO-OpenSource/sgn/pkg/eth"
	"github.com/SeeDAO-OpenSource/sgn/pkg/services"
	"github.com/SeeDAO-OpenSource/sgn/pkg/utils"
	"github.com/ethereum/go-ethereum/ethclient"
)

func AddEthClient(builder *app.AppBuilder) {
	builder.ConfigureServices(func() error {
		var ethOptions = eth.NewEthOptions()
		services.AddValue(ethOptions)
		utils.ViperBind("EthClient", ethOptions)
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
