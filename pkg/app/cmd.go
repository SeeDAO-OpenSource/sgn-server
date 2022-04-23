package app

import (
	"log"

	"github.com/spf13/cobra"
)

type CommandBuilder struct {
	RootCmd       *cobra.Command
	preRunActions []func(cmd *cobra.Command) error
}

func NewCommandBuilder() *CommandBuilder {
	cb := &CommandBuilder{
		preRunActions: []func(cmd *cobra.Command) error{},
	}
	cb.RootCmd = &cobra.Command{
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			cfgfile, err := cmd.Flags().GetString("config")
			if err == nil {
				initConfig(cfgfile)
			}
			for _, action := range cb.preRunActions {
				err := action(cmd)
				if err != nil {
					log.Fatal(err)
				}
			}
			return err
		},
	}
	return cb
}

func (b *CommandBuilder) Build() *cobra.Command {
	b.RootCmd.PersistentFlags().StringP("config", "c", "", "config file (default is ./nftserver.yaml)")
	return b.RootCmd
}

func (b *CommandBuilder) AddCommand(cmd *cobra.Command) {
	b.RootCmd.AddCommand(cmd)
}

func (b *CommandBuilder) PreRun(action func(cmd *cobra.Command) error) {
	b.preRunActions = append(b.preRunActions, action)
}
