package common

import (
	"github.com/SeeDAO-OpenSource/sgn/pkg/app"
	"github.com/SeeDAO-OpenSource/sgn/pkg/mvc"
	"github.com/SeeDAO-OpenSource/sgn/pkg/services"
	"github.com/SeeDAO-OpenSource/sgn/pkg/utils"
)

func AddHttpClient(ac *app.AppBuilder) {
	ac.ConfigureServices(func() error {
		var httpOptions = &mvc.HttpClientOptions{}
		utils.ViperBind("HttpClient", httpOptions)
		services.TryAddValue(httpOptions)
		services.TryAddTransient(func(c *services.Container) *mvc.RequestClient {
			options := services.Get[mvc.HttpClientOptions]()
			client := mvc.NewRequestClient(options)
			return client
		})
		return nil
	})
}
