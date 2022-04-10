package app

import "github.com/spf13/cobra"

type CommandBuilder struct {
	RootCmd *cobra.Command
}

func NewCommandBuilder() *CommandBuilder {
	return &CommandBuilder{
		RootCmd: &cobra.Command{},
	}
}

func (b *CommandBuilder) AddCommand(cmd *cobra.Command) {
	b.RootCmd.AddCommand(cmd)
}

func (b *CommandBuilder) Build() *cobra.Command {
	b.RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.ntfserver.yaml)")
	return b.RootCmd
}
