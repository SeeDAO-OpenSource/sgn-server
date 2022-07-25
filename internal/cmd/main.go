package cmd

import (
	membercmd "github.com/SeeDAO-OpenSource/sgn/internal/cmd/member"
	"github.com/SeeDAO-OpenSource/sgn/internal/cmd/sgn"
	"github.com/SeeDAO-OpenSource/sgn/pkg/app"
	"github.com/spf13/cobra"
)

func AppCommands(ac *app.AppBuilder) {
	ac.CmdBuilder.AddCommand((*cobra.Command)(NewTestCmd()))
	ac.CmdBuilder.AddCommand((*cobra.Command)(NewApiServerCmd(ac)))
	ac.CmdBuilder.AddCommand((*cobra.Command)(NewSgnPullCmd(ac)))
	ac.CmdBuilder.AddCommand((*cobra.Command)(NewConfigCmd()))
	ac.CmdBuilder.AddCommand((*cobra.Command)(NewUpdateCmd()))
	ac.CmdBuilder.AddCommand(membercmd.NewIdentityCmd(ac))
	ac.CmdBuilder.AddCommand(sgn.NewSgnCmd())
}
