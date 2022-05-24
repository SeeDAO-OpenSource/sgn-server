package sgnv1

import (
	"github.com/gin-gonic/gin"
)

func route(sgnctl *SgnController, g *gin.Engine) {
	sgnGroup := g.Group("/api/v1")
	{
		v1 := sgnGroup.Group("sgn")
		{
			v1.GET("", sgnctl.GetOwners)
			v1.GET("image/:token", sgnctl.GetImage)
		}
	}
	g.Static("/app", "./app")
}
