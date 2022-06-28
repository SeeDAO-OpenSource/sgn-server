package memberapi

import (
	"github.com/SeeDAO-OpenSource/sgn/internal/member"
	"github.com/SeeDAO-OpenSource/sgn/pkg/server"
	"github.com/SeeDAO-OpenSource/sgn/pkg/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func MemberApi(builder *server.ServerBuiler) error {
	builder.Configure(func(s *server.Server) error {
		return initRoute(s.G)
	})
	builder.App.ConfigureServices(func() error {
		services.AddTransient(func(c *services.Container) *member.MemberService {
			mongoClient := services.Get[mongo.Client]()
			if mongoClient == nil {
				return nil
			}
			srv, err := member.NewMemberService(mongoClient)
			if err != nil {
				return nil
			}
			return srv
		})
		return nil
	})
	return nil
}
func initRoute(g *gin.Engine) error {
	controller := NewMemberController()
	route(&controller, g)
	return nil
}
