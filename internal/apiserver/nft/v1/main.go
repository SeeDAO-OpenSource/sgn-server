package nftv1

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/waite-lee/nftserver/pkg/app"
	"github.com/waite-lee/nftserver/pkg/server"
)

func InstallNftV1(ac *app.AppContext, server *server.ServerContext) error {
	server.Route(initRoute)
	err := SubscribeTransferLogs()
	return err
}

func initRoute(g *gin.Engine) {
	ntfCtl := newNtfController()
	route(&ntfCtl, g)
}

func SubscribeTransferLogs() error {
	ntfService, err := BuildNtfServiceV1()
	if err == nil {
		address, err := ntfService.GetExistsAddresses()
		if err == nil {
			err1 := ntfService.SubscribeTransferLogs(address)
			if err1 != nil {
				log.Fatal(err1)
			}
		}
	}
	return err
}
