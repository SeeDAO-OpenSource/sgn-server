package nftv1

import (
	"github.com/google/wire"
	"github.com/waite-lee/sgn/internal/apiserver/pkg/erc721"
	"github.com/waite-lee/sgn/internal/common"
	"github.com/waite-lee/sgn/pkg/eth"
)

var NftV1Set = wire.NewSet(
	common.CommonSet,
	erc721.NewErc721Service,
	erc721.GetClient,
	wire.Value(EsOptions),
	eth.GetClient,
	NewMongoErc721TransferLogRepo,
	NewMongoDbNftTokenRepo,
	wire.Struct(new(NftService), "*"),
)
