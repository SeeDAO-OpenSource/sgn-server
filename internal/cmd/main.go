package cmd

import (
	"github.com/google/wire"
	"github.com/spf13/cobra"
	"github.com/waite-lee/nftserver/pkg/app"
)

var CmdSet = wire.NewSet(
	wire.Struct(new(AppCommands), "*"),
	NewTestCmd,
	NewApiServerCmd,
	NewNftPullCmd,
	NewConfigCmd,
	NewUpdateCmd,
)

type AppCommands struct {
	Test      *TestCmd
	ApiServer *ApiServerCmd
	NftPull   *NftPullCmd
	Config    *ConfigCmd
	Update    *UpdateCmd
}

func (ac *AppCommands) Build() {

}

func AddCommands(ac *app.AppContext) {
	cmds := BuildCommands(ac)
	ac.CmdBuilder.AddCommand((*cobra.Command)(cmds.Test))
	ac.CmdBuilder.AddCommand((*cobra.Command)(cmds.ApiServer))
	ac.CmdBuilder.AddCommand((*cobra.Command)(cmds.NftPull))
	ac.CmdBuilder.AddCommand((*cobra.Command)(cmds.Config))
	ac.CmdBuilder.AddCommand((*cobra.Command)(cmds.Update))
}
