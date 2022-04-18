package server

import "github.com/gin-gonic/gin"

type ServerContext struct {
	GEngine *gin.Engine
	router  []RouteBuildFunc
}

func NewServerContext() *ServerContext {
	return &ServerContext{
		GEngine: gin.Default(),
		router:  make([]RouteBuildFunc, 0, 0),
	}
}

func (ac *ServerContext) Route(buildFunc RouteBuildFunc) {
	ac.router = append(ac.router, buildFunc)
}

func (ac *ServerContext) Init() {
	if ac.router != nil {
		for _, r := range ac.router {
			r(ac.GEngine)
		}
	}

}
