package blobv1

import (
	"github.com/SeeDAO-OpenSource/sgn/internal/common"
	"github.com/google/wire"
)

var BlobV1Set = wire.NewSet(
	common.CommonSet,
	NewBlobService,
)
