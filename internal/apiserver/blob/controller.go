package blob

import (
	"errors"
	"net/http"
	"strings"

	"github.com/SeeDAO-OpenSource/sgn/pkg/blob"
	"github.com/SeeDAO-OpenSource/sgn/pkg/mvc"
	"github.com/SeeDAO-OpenSource/sgn/pkg/services"
	"github.com/gin-gonic/gin"
)

type BlobController struct {
}

func newBlobController() BlobController {
	return BlobController{}
}

// @Summary Get blob data
// @Schemes
// @Description
// @Tags example
// @Produce octet-stream
// @Success 200
// @Router /api/blob/{key} [get]
// @Param key path string true "key"
func (c *BlobController) Get(ctx *gin.Context) {
	key := ctx.Param("key")
	if key == "" {
		mvc.Error(ctx, errors.New("key is required"))
		return
	}
	key = strings.TrimLeft(key, "/")
	service := services.Get[BlobService]()
	if service == nil {
		mvc.Error(ctx, errors.New("blob service is nil"))
		return
	}
	reader, err := service.Get(key, parseProcess(ctx))
	if err != nil {
		mvc.Error(ctx, err)
		return
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
