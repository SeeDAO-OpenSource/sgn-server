package nftv1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/waite-lee/sgn/pkg/blob"
	"github.com/waite-lee/sgn/pkg/mvc"
)

type NftController struct {
}

func newNftController() NftController {
	return NftController{}
}

func (c *NftController) GetOwners(ctx *gin.Context) {
	page, pageSize := mvc.PageQuery(ctx)
	srv, err := BuildNftServiceV1()
	if err != nil {
		mvc.Error(ctx, err)
		return
	}
	data, err := srv.GetOwners("0x23fDA8a873e9E46Dbe51c78754dddccFbC41CFE1", page, pageSize)
	if err != nil {
		mvc.Error(ctx, err)
		return
	}
	mvc.Ok(ctx, data)
}

func (c *NftController) GetImage(ctx *gin.Context) {
	token := ctx.Param("token")
	service, err := BuildNftServiceV1()
	reader, err := service.GetTokenImage(token, "0x23fDA8a873e9E46Dbe51c78754dddccFbC41CFE1", parseProcess(ctx))
	if err != nil {
		mvc.Error(ctx, err)
		return
	}
	if reader == nil {
		mvc.Error(ctx, errors.New("image not found"))
	}
	contentType := "application/octet-stream"
	if len(reader.Content) >= 512 {
		contentType = http.DetectContentType(reader.Content[:512])
	}
	ctx.Data(200, contentType, reader.Content)
}

func parseProcess(ctx *gin.Context) *blob.Process {
	process := blob.Process{}
	process.Width = mvc.QueryInt(ctx, "w")
	process.Height = mvc.QueryInt(ctx, "h")
	return &process
}
