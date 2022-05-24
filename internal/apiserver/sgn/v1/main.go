package sgnv1

import (
	"github.com/SeeDAO-OpenSource/sgn/internal/apiserver/pkg/erc721"
	"github.com/SeeDAO-OpenSource/sgn/pkg/server"
	"github.com/gin-gonic/gin"
)

var EsOptions = &erc721.EtherScanOptions{
	BaseURL: "https://api.etherscan.io/api?",
}

func AddSgnV1(builder *server.ServerBuiler) error {
	builder.Configure(func(s *server.Server) error {
		return initRoute(s.G)
	})
	builder.AppBuilder.BindOptions("EtherScan", EsOptions)
	return nil
}

func initRoute(g *gin.Engine) error {
	sgnCtl := newSgnController()
	route(&sgnCtl, g)
	SubscribeTransferLogs()
	return nil
}

func SubscribeTransferLogs() error {
	sgnService, err := BuildSgnServiceV1()
	if err != nil {
		return err
	}
	addresses, err := sgnService.GetExistsAddresses()
	if err != nil {
		return err
	}
	if len(addresses) > 0 {
		err = sgnService.SubscribeTransferLogs(addresses)
		if err != nil {
			return err
		}
	}
	return nil
}
