//go:build wireinject
// +build wireinject

package nftv1

import (
	"github.com/google/wire"
)

func BuildNftServiceV1() (*NftService, error) {
	wire.Build(
		NftV1Set,
	)
	return &NftService{}, nil
}
