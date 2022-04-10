package nftv1

import (
	"github.com/gin-gonic/gin"
	nnet "github.com/waite-lee/nftserver/pkg/net"
)

type NftController struct {
}

func newNtfController() NftController {
	return NftController{}
}

func (c *NftController) Get(ctx *gin.Context) {
	nnet.Ok(ctx, "hello NFT")
}
