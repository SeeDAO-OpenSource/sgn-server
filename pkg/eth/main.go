package eth

import (
	"sync"

	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	client *ethclient.Client
	once   sync.Once
)

func GetClient(options *EthOptions) (*ethclient.Client, error) {
	var err error = nil
	once.Do(func() {
		client, err = ethclient.Dial(options.DailUrl)
	})
	return client, err
}
