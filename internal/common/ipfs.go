package common

import (
	"github.com/SeeDAO-OpenSource/sgn/pkg/app"
	"github.com/SeeDAO-OpenSource/sgn/pkg/ipfs"
	"github.com/SeeDAO-OpenSource/sgn/pkg/mvc"
	"github.com/SeeDAO-OpenSource/sgn/pkg/services"
	"github.com/SeeDAO-OpenSource/sgn/pkg/utils"
)

func AddIpfsClient(builder *app.AppBuilder) {
	builder.ConfigureServices(func() error {
		var ipfsOptions = &ipfs.IpfsOptions{
			BaseURL: "https://ipfs.io/ipfs/",
		}
		utils.ViperBind("Ipfs", ipfsOptions)
		services.AddValue(ipfsOptions)
		services.AddTransient(func(c *services.Container) *ipfs.IpfsClient {
			options := services.Get[ipfs.IpfsOptions]()
			requestClient := services.Get[mvc.RequestClient]()
			if requestClient == nil || options == nil {
				return nil
			}
			client, err := ipfs.GetClient(requestClient, options)
			if err != nil {
				return nil
			}
			return client
		})
		return nil
	})
}
