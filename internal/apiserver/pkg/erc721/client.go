package erc721

import (
	"errors"
	"sync"

	"github.com/nanmu42/etherscan-api"
	"github.com/waite-lee/nftserver/pkg/mvc"
)

var (
	client *etherscan.Client
	once   sync.Once
)

func GetClient(options *EtherScanOptions, httpOptions *mvc.HttpClientOptions) (*etherscan.Client, error) {
	if options.ApiKey == "" {
		return nil, errors.New("未配置 EtherScan Api Key")
	}
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
