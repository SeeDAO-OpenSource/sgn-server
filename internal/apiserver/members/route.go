package memebersv1

import "github.com/gin-gonic/gin"

func route(ctl *IdentityController, g *gin.Engine) {
	sgnGroup := g.Group("/api/v1")
	{
		v1 := sgnGroup.Group("identity")
		{
			v1.GET("", ctl.GetList)
			v1.POST("", ctl.Insert)
			v1.GET(":address", ctl.GetByAddress)
			v1.GET("range/:addresses", ctl.GetByAddresses)
			v1.PUT(":address", ctl.Update)
			v1.DELETE(":address", ctl.Delete)
		}
	}
}
