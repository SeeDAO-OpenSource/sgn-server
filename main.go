package main

import (
	"log"

	"github.com/waite-lee/sgn/internal/cmd"
	"github.com/waite-lee/sgn/pkg/app"
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
	builder.Version(GetVersion().Version)
	cmd.AddCommands(builder)
	builder.Info("sgn", "提供SeeDao SGN 相关功能", "提供SeeDao SGN 相关功能")
	return builder.Build()
}
