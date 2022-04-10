package nftv1

import (
	"github.com/gin-gonic/gin"
	"github.com/waite-lee/nftserver/pkg/mvc"
)

type NftController struct {
}

func newNtfController() NftController {
	return NftController{}
}

func (c *NftController) GetList(ctx *gin.Context) {
	mvc.Ok(ctx, "hello NFT")
}
