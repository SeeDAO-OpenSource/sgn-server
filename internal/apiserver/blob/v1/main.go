package blobv1

import (
	"github.com/gin-gonic/gin"
	"github.com/waite-lee/sgn/pkg/app"
	"github.com/waite-lee/sgn/pkg/server"
)

func InstallBlobV1(ac *app.AppContext, server *server.ServerContext) error {
	server.Route(initRoute)
	return nil
}

func initRoute(g *gin.Engine) {
	nftCtl := newBlobController()
	route(&nftCtl, g)
}
