package blobv1

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/waite-lee/nftserver/pkg/mvc"
)

type BlobController struct {
}

func newBlobController() BlobController {
	return BlobController{}
}

func (c *BlobController) Get(ctx *gin.Context) {
	key := ctx.Param("key")
	if key == "" {
		mvc.Error(ctx, errors.New("key is required"))
		return
	}
	key = strings.TrimLeft(key, "/")
	service := BuildBlobServiceV1()
	reader, err := service.Get(key)
	if err != nil {
		mvc.Error(ctx, err)
		return
	}
	contentType := http.DetectContentType(reader.Content[:512])
	ctx.Data(200, contentType, reader.Content)
}
