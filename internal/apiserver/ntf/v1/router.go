package nftv1

import (
	"github.com/gin-gonic/gin"
)

func route(ntfctl *NftController, g *gin.Engine) {
	ntfGroup := g.Group("/api/v1")
	{
		v1 := ntfGroup.Group("nft")
		{
			v1.GET("", ntfctl.GetList)
		}
	}
}
