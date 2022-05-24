package erc721

import (
	"sync"

	"github.com/SeeDAO-OpenSource/sgn/pkg/mvc"
	"github.com/nanmu42/etherscan-api"
)

var (
	client *etherscan.Client
	once   sync.Once
)

func GetClient(options *EtherScanOptions, httpOptions *mvc.HttpClientOptions) (*etherscan.Client, error) {
	once.Do(func() {
		httpClient := mvc.NewHttpClient(httpOptions)
		client = etherscan.NewCustomized(etherscan.Customization{
			Key:     options.ApiKey,
			BaseURL: options.BaseURL,
			Client:  httpClient,
		})
	})
	return client, nil
}
