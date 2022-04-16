package mvc

import (
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
)

var (
	client *resty.Client
	once   sync.Once
)

type HttpClientOptions struct {
	ProxyURL string
}

type RequestClient struct {
	options *HttpClientOptions
}

func NewRequestClient(hoptions *HttpClientOptions) *RequestClient {
	return &RequestClient{
		options: hoptions,
	}
}

func NewHttpClient(hoptions *HttpClientOptions) *http.Client {
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
	if hoptions.ProxyURL != "" {
		proxyUrl, _ := url.Parse(hoptions.ProxyURL)
		transport.Proxy = http.ProxyURL(proxyUrl)
	}
	httpClient.Transport = transport
	return httpClient
}

func getDefaultClient(hoptions *HttpClientOptions) *resty.Client {
	once.Do(func() {
		hc := NewHttpClient(hoptions)
		client = resty.NewWithClient(hc)
	})
	return client
}

func (c *RequestClient) Get(url string) ([]byte, error) {
	hc := getDefaultClient(c.options)
	resp, err := hc.R().
		EnableTrace().
		Get(url)
	if err != nil {
		return []byte{}, err
	}
	return resp.Body(), nil
}

func (c *RequestClient) GetString(url string) (string, error) {
	hc := getDefaultClient(c.options)
	resp, err := hc.R().
		EnableTrace().
		Get(url)
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}
