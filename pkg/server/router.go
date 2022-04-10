package server

import "github.com/gin-gonic/gin"

type RouteBuildFunc func(g *gin.Engine)
