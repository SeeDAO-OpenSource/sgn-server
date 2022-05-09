//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/google/wire"
	"github.com/waite-lee/sgn/internal/apiserver"
	"github.com/waite-lee/sgn/pkg/app"
)

func BuildCommands(ac *app.AppContext) AppCommands {
	wire.Build(
		apiserver.ApiServerSet,
		CmdSet,
	)
	return AppCommands{}
}
