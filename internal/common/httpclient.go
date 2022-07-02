package common

import (
	"github.com/SeeDAO-OpenSource/sgn/pkg/app"
	"github.com/SeeDAO-OpenSource/sgn/pkg/di"
	"github.com/SeeDAO-OpenSource/sgn/pkg/mvc"
	"github.com/SeeDAO-OpenSource/sgn/pkg/utils"
)

func AddHttpClient(ac *app.AppBuilder) {
	ac.ConfigureServices(func() error {
		var httpOptions = &mvc.HttpClientOptions{}
		utils.ViperBind("HttpClient", httpOptions)
		di.TryAddValue(httpOptions)
		di.TryAddTransient(func(c *di.Container) *mvc.RequestClient {
			options := di.Get[mvc.HttpClientOptions]()
			client := mvc.NewRequestClient(options)
			return client
		})
		return nil
	})
}
