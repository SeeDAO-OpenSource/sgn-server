package blobv1

import (
	"github.com/google/wire"
	"github.com/waite-lee/sgn/internal/common"
)

var BlobV1Set = wire.NewSet(
	common.CommonSet,
	NewBlobService,
)
