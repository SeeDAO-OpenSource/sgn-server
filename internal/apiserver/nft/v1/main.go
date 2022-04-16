package nftv1

import (
	"github.com/gin-gonic/gin"
	"github.com/waite-lee/nftserver/pkg/app"
	"github.com/waite-lee/nftserver/pkg/server"
)

func InstallNftV1(ac *app.AppContext, server *server.ServerContext) {
	server.Route(initRoute)
}

func initRoute(g *gin.Engine) {
	ntfCtl := newNtfController()
	route(&ntfCtl, g)
}
