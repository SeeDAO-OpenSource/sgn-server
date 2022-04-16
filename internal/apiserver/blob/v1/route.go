package blobv1

import (
	"github.com/gin-gonic/gin"
)

func route(ntfctl *BlobController, g *gin.Engine) {
	ntfGroup := g.Group("/api/v1")
	{
		v1 := ntfGroup.Group("blob")
		{
			v1.Any("*key", ntfctl.Get)
		}
	}
}
