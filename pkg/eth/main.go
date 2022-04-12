package eth

import (
	"sync"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/wire"
)

var EthSet = wire.NewSet(
	GetClient,
	wire.Value(Options),
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
