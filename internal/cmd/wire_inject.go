//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/google/wire"
	"github.com/waite-lee/sgn/pkg/app"
)

func BuildCommands(buider *app.AppBuilder) AppCommands {
	wire.Build(
		CmdSet,
	)
	return AppCommands{}
}
