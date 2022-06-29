package main

import (
	"log"

	"github.com/SeeDAO-OpenSource/sgn/internal/cmd"
	"github.com/SeeDAO-OpenSource/sgn/internal/common"
	"github.com/SeeDAO-OpenSource/sgn/pkg/app"
)

func main() {
	app, err := buildApp()
	if err != nil {
		log.Fatal("初始化失败:" + err.Error())
	}
	app.Run()
}

func buildApp() (*app.App, error) {
	builder := app.NewAppBuilder()
	builder.Version(GetVersion().Version).
		Info("sgn", "提供SeeDao SGN 相关功能", "提供SeeDao SGN 相关功能").
		Use(common.CommonServices).
		Use(cmd.AppCommands)
	return builder.Build()
}
