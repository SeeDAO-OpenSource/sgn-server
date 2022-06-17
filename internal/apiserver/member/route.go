package memberapi

import "github.com/gin-gonic/gin"

func route(ctl *MemberController, g *gin.Engine) {
	sgnGroup := g.Group("/api")
	{
		v1 := sgnGroup.Group("members")
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
