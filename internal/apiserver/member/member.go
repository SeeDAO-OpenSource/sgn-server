package memberapi

import (
	"github.com/SeeDAO-OpenSource/sgn/pkg/server"
	"github.com/gin-gonic/gin"
)

func AddMember(builder *server.ServerBuiler) error {
	builder.Configure(func(s *server.Server) error {
		return initRoute(s.G)
	})
	return nil
}
func initRoute(g *gin.Engine) error {
	controller := NewMemberController()
	route(&controller, g)
	return nil
}
