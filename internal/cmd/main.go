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
)

type AppCommands struct {
	Test      *TestCmd
	ApiServer *ApiServerCmd
	NtfPull   *NftPullCmd
}

func (ac *AppCommands) Build() {

}

func InstallCommands(ac *app.AppContext) {
	cmds := BuildCommands(ac)
	ac.CmdBuilder.AddCommand((*cobra.Command)(cmds.Test))
	ac.CmdBuilder.AddCommand((*cobra.Command)(cmds.ApiServer))
	ac.CmdBuilder.AddCommand((*cobra.Command)(cmds.NtfPull))
}
