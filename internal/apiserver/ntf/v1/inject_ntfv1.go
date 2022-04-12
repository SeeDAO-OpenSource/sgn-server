//go:build wireinject
// +build wireinject

package nftv1

import (
	"github.com/google/wire"
	"github.com/waite-lee/nftserver/internal/apiserver/ntf/erc721"
	"github.com/waite-lee/nftserver/pkg/eth"
)

func BuildNtfServiceV1() (*NftService, error) {
	wire.Build(
		erc721.NewErc721Service,
		erc721.GetClient,
		wire.Value(erc721.EsOptions),
		eth.EthSet,
		NewNtfService,
	)
	return &NftService{}, nil
}
