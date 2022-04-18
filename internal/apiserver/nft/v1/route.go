package nftv1

import (
	"github.com/gin-gonic/gin"
)

func route(nftctl *NftController, g *gin.Engine) {
	nftGroup := g.Group("/api/v1")
	{
		v1 := nftGroup.Group("nft")
		{
			v1.GET(":nftaddr", nftctl.GetOwners)
		}
	}

	g.Static("/app", "./app")
}
