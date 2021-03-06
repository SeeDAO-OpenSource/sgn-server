package sgn

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/SeeDAO-OpenSource/sgn/pkg/blob"
	"github.com/SeeDAO-OpenSource/sgn/pkg/di"
	"github.com/SeeDAO-OpenSource/sgn/pkg/mvc"
	"github.com/gin-gonic/gin"
)

type SgnController struct {
}

func newSgnController() SgnController {
	return SgnController{}
}

// @Summary 获取sgn信息列表
// @Schemes
// @Description 按照交易时间进行降序排序
// @Tags sgn
// @Accept json
// @Produce json
// @Success 200
// @Router /api/sgn [get]
// @Param page query int false "页码"
// @Param pageSize query int false "每页数量"
func (c *SgnController) GetOwners(ctx *gin.Context) {
	skip := mvc.QueryInt(ctx, "skip")
	limit := mvc.QueryInt(ctx, "limit")
	srv := di.Get[SgnService]()
	if srv == nil {
		mvc.Error(ctx, errors.New("sgn service is nil"))
		return
	}
	data, err := srv.GetOwners("0x23fDA8a873e9E46Dbe51c78754dddccFbC41CFE1", int64(skip), int64(limit))
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
	service := di.Get[SgnService]()
	if service == nil {
		mvc.Error(ctx, errors.New("sgn service is nil"))
		return
	}
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
