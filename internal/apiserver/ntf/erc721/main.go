package erc721

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/metachris/eth-go-bindings/erc721"
	"github.com/nanmu42/etherscan-api"
)

type Erc721Service struct {
	options *EtherScanOptions
	client  *etherscan.Client
}

type TokenInfo struct {
	Name     string
	TokenId  *big.Int
	TokenURI string
}

func NewErc721Service(client *etherscan.Client, options *EtherScanOptions) *Erc721Service {
	service := &Erc721Service{
		options: options,
		client:  client,
	}
	return service
}

func (srv *Erc721Service) GetTransferLogs(address *string, page int, pageSize int) ([]etherscan.ERC721Transfer, error) {
	return srv.client.ERC721Transfers(address, nil, nil, nil, page, pageSize, true)
}

func (srv *Erc721Service) GetToken(ethClient *ethclient.Client, address *string, tokenId *big.Int) (*TokenInfo, error) {
	contrctAddr := common.HexToAddress(*address)
	token, err := erc721.NewErc721(contrctAddr, ethClient)
	if err != nil {
		return nil, err
	}
	info := TokenInfo{
		TokenId: tokenId,
	}
	name, err := token.Name(nil)
	info.Name = name
	tokenURI, err := token.TokenURI(nil, tokenId)
	info.TokenURI = tokenURI
	return &info, nil
}
