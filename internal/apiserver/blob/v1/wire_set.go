package blobv1

import (
	"github.com/google/wire"
	"github.com/waite-lee/nftserver/internal/common"
)

var BlobV1Set = wire.NewSet(
	common.CommonSet,
	NewBlobService,
)
