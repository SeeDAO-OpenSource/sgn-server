package swagger

import (
	"github.com/SeeDAO-OpenSource/sgn/internal/apiserver/docs"
	"github.com/SeeDAO-OpenSource/sgn/pkg/server"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SwaggerDoc(builder *server.ServerBuiler) error {
	docs.SwaggerInfo.BasePath = "/api"
	builder.Configure(func(s *server.Server) error {
		s.G.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
		return nil
	})
	return nil
}
