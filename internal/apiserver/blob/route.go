package blob

import (
	"github.com/gin-gonic/gin"
)

func route(blobCtroller *BlobController, g *gin.Engine) {
	sgnGroup := g.Group("/api")
	{
		v1 := sgnGroup.Group("blob")
		{
			v1.GET("*key", func(ctx *gin.Context) {
				blobCtroller.Get(ctx)
			})
		}
	}
}
