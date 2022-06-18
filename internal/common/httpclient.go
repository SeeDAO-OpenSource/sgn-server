package common

import (
	"github.com/SeeDAO-OpenSource/sgn/pkg/app"
	"github.com/SeeDAO-OpenSource/sgn/pkg/mvc"
	"github.com/SeeDAO-OpenSource/sgn/pkg/services"
)

func AddHttpClient(ac *app.AppBuilder) {
	var httpOptions = &mvc.HttpClientOptions{}
	ac.BindOptions("HttpClient", httpOptions)
	ac.ConfigureServices(func() error {
		services.TryAddValue(httpOptions)
		services.TryAddTransient(func(c *services.Container) *mvc.RequestClient {
			options := services.Get[mvc.HttpClientOptions]()
			client := mvc.NewRequestClient(options)
			return client
		})
		return nil
	})
}
