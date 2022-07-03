package common

import (
	"github.com/SeeDAO-OpenSource/sgn/pkg/app"
	"github.com/SeeDAO-OpenSource/sgn/pkg/di"
	"github.com/SeeDAO-OpenSource/sgn/pkg/ipfs"
	"github.com/SeeDAO-OpenSource/sgn/pkg/mvc"
	"github.com/SeeDAO-OpenSource/sgn/pkg/utils"
)

func AddIpfsClient(builder *app.AppBuilder) {
	builder.ConfigureServices(func() error {
		var ipfsOptions = &ipfs.IpfsOptions{
			BaseURL: "https://ipfs.io/ipfs/",
		}
		utils.ViperBind("Ipfs", ipfsOptions)
		di.AddValue(ipfsOptions)
		di.AddTransient(func(c *di.Container) *ipfs.IpfsClient {
			options := di.Get[ipfs.IpfsOptions]()
			requestClient := di.Get[mvc.RequestClient]()
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
