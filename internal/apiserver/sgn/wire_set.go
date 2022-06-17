package sgn

import (
	"github.com/SeeDAO-OpenSource/sgn/internal/apiserver/pkg/erc721"
	"github.com/SeeDAO-OpenSource/sgn/internal/common"
	"github.com/SeeDAO-OpenSource/sgn/pkg/eth"
	"github.com/google/wire"
)

var sgnSet = wire.NewSet(
	common.CommonSet,
	erc721.NewErc721Service,
	erc721.GetClient,
	wire.Value(EsOptions),
	eth.GetClient,
	NewMongoErc721TransferLogRepo,
	NewMongoDbSgnTokenRepo,
	wire.Struct(new(SgnService), "*"),
)
