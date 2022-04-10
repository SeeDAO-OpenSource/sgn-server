package erc721

import (
	"errors"

	"github.com/nanmu42/etherscan-api"
)

type Erc721Service struct {
	options *EtherScanOptions
}

func NewErc721Service(options *EtherScanOptions) *Erc721Service {
	return &Erc721Service{
		options: options,
	}
}

func (srv *Erc721Service) getTransferLogs(address *string, page int, pageSize int) ([]etherscan.ERC721Transfer, error) {
	if srv.options.ApiKey == "" {
		return []etherscan.ERC721Transfer{}, errors.New("未配置 EtherScan Api Key")
	}
	esClient := etherscan.NewCustomized(etherscan.Customization{
		Key:     srv.options.ApiKey,
		BaseURL: srv.options.BaseURL,
	})
	return esClient.ERC721Transfers(address, nil, nil, nil, page, pageSize, true)
}
