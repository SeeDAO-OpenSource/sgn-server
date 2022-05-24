//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/SeeDAO-OpenSource/sgn/pkg/app"
	"github.com/google/wire"
)

func BuildCommands(buider *app.AppBuilder) AppCommands {
	wire.Build(
		CmdSet,
	)
	return AppCommands{}
}
