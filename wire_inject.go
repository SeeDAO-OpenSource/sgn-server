//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/waite-lee/sgn/pkg/app"
)

func BuildAppContext() (*app.AppContext, error) {
	wire.Build(
		app.NewAppContext,
		app.NewCommandBuilder,
	)
	return &app.AppContext{}, nil
}
