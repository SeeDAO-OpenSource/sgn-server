package ipfs

import (
	"strings"
	"sync"

	"github.com/SeeDAO-OpenSource/sgn/pkg/mvc"
)

type IpfsOptions struct {
	BaseURL string
}

type IpfsClient struct {
	options       *IpfsOptions
	requestClient *mvc.RequestClient
}

var (
	client *IpfsClient
	once   sync.Once
)

func GetClient(requestClient *mvc.RequestClient, moptions *IpfsOptions) (*IpfsClient, error) {
	var err error = nil
	once.Do(func() {
		client = &IpfsClient{
			options:       moptions,
			requestClient: requestClient,
		}
	})
	if err != nil {
		return nil, err
	}
	return client, err
}

func (c *IpfsClient) GetString(uri string) (string, error) {
	url := c.fullUrl(uri)
	return c.requestClient.GetString(url)
}

func (c *IpfsClient) GetContent(uri string) ([]byte, error) {
	url := c.fullUrl(uri)
	content, err := c.requestClient.Get(url, true)
	return content, err
}

func (c *IpfsClient) fullUrl(uri string) string {
	return c.options.BaseURL + c.TrimKey(uri)
}

func (c *IpfsClient) TrimKey(uri string) string {
	uri = strings.TrimLeft(uri, "ipfs://")
	uri = strings.TrimLeft(uri, "ipfs/")
	return uri
}
