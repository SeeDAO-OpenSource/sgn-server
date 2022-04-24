package main

import (
	"log"

	"github.com/waite-lee/nftserver/internal/cmd"
	"github.com/waite-lee/nftserver/internal/common"
	"github.com/waite-lee/nftserver/pkg/app"
)

func main() {
	app, err := buildApp()
	if err != nil {
		log.Fatal("初始化失败:" + err.Error())
	}
	app.Run()
}

func buildApp() (app.App, error) {
	appContext, err := BuildAppContext()
	appContext.Version(GetVersion().Version)
	common.AddCommonOptions(appContext)
	cmd.AddCommands(appContext)
	appContext.RootCmd("nftserver", "提供NFT相关功能")
	app := appContext.Build()
	return app, err
}
