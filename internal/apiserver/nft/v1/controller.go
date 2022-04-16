package nftv1

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/waite-lee/nftserver/pkg/mvc"
)

type NftController struct {
}

func newNtfController() NftController {
	return NftController{}
}

func (c *NftController) GetOwners(ctx *gin.Context) {
	addr := ctx.Param("ntfaddr")
	if addr == "" {
		mvc.Error(ctx, errors.New("ntfaddr is required"))
		return
	}
	page, pageSize := mvc.PageQuery(ctx)
	srv, err := BuildNtfServiceV1()
	if err != nil {
		mvc.Error(ctx, err)
		return
	}
	data, err := srv.GetOwners(addr, page, pageSize)
	if err != nil {
		mvc.Error(ctx, err)
		return
	}
	mvc.Ok(ctx, data)
}
