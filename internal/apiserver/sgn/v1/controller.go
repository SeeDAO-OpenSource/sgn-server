package sgnv1

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/SeeDAO-OpenSource/sgn/pkg/blob"
	"github.com/SeeDAO-OpenSource/sgn/pkg/mvc"
	"github.com/gin-gonic/gin"
)

type SgnController struct {
}

func newSgnController() SgnController {
	return SgnController{}
}

func (c *SgnController) GetOwners(ctx *gin.Context) {
	page, pageSize := mvc.PageQuery(ctx)
	srv, err := BuildSgnServiceV1()
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

func (c *SgnController) GetImage(ctx *gin.Context) {
	tokenStr := ctx.Param("token")
	token, err := strconv.ParseInt(tokenStr, 10, 64)
	if err != nil {
		mvc.Error(ctx, err)
		return
	}
	service, err := BuildSgnServiceV1()
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
