//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/waite-lee/nftserver/pkg/app"
)

func BuildAppContext() (*app.AppContext, error) {
	wire.Build(
		app.NewCommandBuilder,
		wire.Struct(new(app.AppContext), "*"),
	)
	return &app.AppContext{}, nil
}
