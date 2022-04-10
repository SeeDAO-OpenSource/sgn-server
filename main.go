package main

import (
	"log"

	"github.com/waite-lee/nftserver/internal/cmd"
)

func main() {
	appContext, err := BuildAppContext()
	cmd.InstallCommands(appContext)
	appContext.AppInfo("NFT服务", "提供NFT相关基础服务")
	app := appContext.Build()
	if err != nil {
		log.Fatal("初始化失败:" + err.Error())
	}
	app.Run()
}
