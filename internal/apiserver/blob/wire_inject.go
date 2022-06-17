//go:build wireinject
// +build wireinject

package blob

import (
	"github.com/google/wire"
)

func BuildBlobServiceV1() *BlobService {
	wire.Build(
		blobSet,
	)
	return &BlobService{}
}
