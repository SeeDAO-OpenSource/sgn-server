package blob

import (
	"github.com/SeeDAO-OpenSource/sgn/internal/common"
	"github.com/google/wire"
)

var blobSet = wire.NewSet(
	common.CommonSet,
	NewBlobService,
)
