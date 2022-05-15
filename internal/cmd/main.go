package cmd

import (
	"github.com/google/wire"
	"github.com/spf13/cobra"
	"github.com/waite-lee/sgn/internal/common"
	"github.com/waite-lee/sgn/pkg/app"
)

var CmdSet = wire.NewSet(
	wire.Struct(new(AppCommands), "*"),
	NewTestCmd,
	NewApiServerCmd,
	NewNftPullCmd,
	NewConfigCmd,
	NewUpdateCmd,
	common.CommonSet,
)

type AppCommands struct {
	Test      *TestCmd
	ApiServer *ApiServerCmd
	NftPull   *NftPullCmd
	Config    *ConfigCmd
	Update    *UpdateCmd
}

func AddCommands(ac *app.AppBuilder) {
	cmds := BuildCommands(ac)
	ac.CmdBuilder.AddCommand((*cobra.Command)(cmds.Test))
	ac.CmdBuilder.AddCommand((*cobra.Command)(cmds.ApiServer))
	ac.CmdBuilder.AddCommand((*cobra.Command)(cmds.NftPull))
	ac.CmdBuilder.AddCommand((*cobra.Command)(cmds.Config))
	ac.CmdBuilder.AddCommand((*cobra.Command)(cmds.Update))
}
