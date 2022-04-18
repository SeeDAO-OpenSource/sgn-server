package nftv1

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/waite-lee/nftserver/pkg/app"
	"github.com/waite-lee/nftserver/pkg/server"
)

func InstallNftV1(ac *app.AppContext, server *server.ServerContext) error {
	server.Route(initRoute)
	ac.CmdBuilder.PreRun(func(cmd *cobra.Command) error {
		bindOptions()
		err := SubscribeTransferLogs()
		return err
	})
	return nil
}

func initRoute(g *gin.Engine) {
	nftCtl := newNftController()
	route(&nftCtl, g)
}

func SubscribeTransferLogs() error {
	nftService, err := BuildNftServiceV1()
	if err != nil {
		return err
	}
	address, err := nftService.GetExistsAddresses()
	if err != nil {
		return err
	}
	err = nftService.SubscribeTransferLogs(address)
	if err != nil {
		return err
	}
	return err
}

func bindOptions() {
	viper.UnmarshalKey("EtherScan", EsOptions)
}
