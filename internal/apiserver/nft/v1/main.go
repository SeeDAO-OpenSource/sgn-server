package nftv1

import (
	"github.com/gin-gonic/gin"
	"github.com/waite-lee/sgn/internal/apiserver/pkg/erc721"
	"github.com/waite-lee/sgn/pkg/server"
)

var EsOptions = &erc721.EtherScanOptions{
	BaseURL: "https://api.etherscan.io/api?",
}

func AddNftV1(builder *server.ServerBuiler) error {
	builder.Configure(func(s *server.Server) error {
		return initRoute(s.G)
	})
	builder.AppBuilder.BindOptions("EtherScan", EsOptions)
	return nil
}

func initRoute(g *gin.Engine) error {
	nftCtl := newNftController()
	route(&nftCtl, g)
	SubscribeTransferLogs()
	return nil
}

func SubscribeTransferLogs() error {
	nftService, err := BuildNftServiceV1()
	if err != nil {
		return err
	}
	addresses, err := nftService.GetExistsAddresses()
	if err != nil {
		return err
	}
	if len(addresses) > 0 {
		err = nftService.SubscribeTransferLogs(addresses)
		if err != nil {
			return err
		}
	}
	return nil
}
