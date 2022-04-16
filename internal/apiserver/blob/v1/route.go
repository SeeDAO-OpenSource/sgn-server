package blobv1

import (
	"github.com/gin-gonic/gin"
)

func route(nftctl *BlobController, g *gin.Engine) {
	nftGroup := g.Group("/api/v1")
	{
		v1 := nftGroup.Group("blob")
		{
			v1.Any("*key", nftctl.Get)
		}
	}
}
