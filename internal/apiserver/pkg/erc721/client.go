package erc721

import (
	"errors"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/nanmu42/etherscan-api"
)

var (
	client *etherscan.Client
	once   sync.Once
)

func GetClient(options *EtherScanOptions) (*etherscan.Client, error) {
	if options.ApiKey == "" {
		return nil, errors.New("未配置 EtherScan Api Key")
	}
	once.Do(func() {
		httpClient := &http.Client{}
		transport := &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		}
		if options.Proxy != "" {
			proxyUrl, _ := url.Parse(options.Proxy)
			transport.Proxy = http.ProxyURL(proxyUrl)
		}
		httpClient.Transport = transport
		client = etherscan.NewCustomized(etherscan.Customization{
			Key:     options.ApiKey,
			BaseURL: options.BaseURL,
			Client:  httpClient,
		})
	})
	return client, nil
}
