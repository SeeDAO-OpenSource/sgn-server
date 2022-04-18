//go:build wireinject
// +build wireinject

package blobv1

import (
	"github.com/google/wire"
)

func BuildBlobServiceV1() *BlobService {
	wire.Build(
		BlobV1Set,
	)
	return &BlobService{}
}
