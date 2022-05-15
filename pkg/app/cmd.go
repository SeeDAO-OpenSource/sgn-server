package app

import (
	"github.com/spf13/cobra"
)

type CommandBuilder struct {
	RootCmd *cobra.Command
}

func NewCommandBuilder() *CommandBuilder {
	cb := &CommandBuilder{
		RootCmd: &cobra.Command{},
	}
	return cb
}

func (cb *CommandBuilder) Build(ac *AppContext) (*cobra.Command, error) {
	cb.RootCmd.Use = ac.Name
	cb.RootCmd.Short = ac.Short
	cb.RootCmd.Long = ac.Description
	cb.RootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		return preRunApp(cmd, ac)
	}
	cb.RootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runApp(ac)
	}
	cb.RootCmd.PersistentPostRunE = func(cmd *cobra.Command, args []string) error {
		return postRunApp(ac)
	}
	err := configureApp(cb.RootCmd, ac)
	if err != nil {
		return nil, err
	}
	return cb.RootCmd, nil
}

func (b *CommandBuilder) AddCommand(cmd *cobra.Command) {
	b.RootCmd.AddCommand(cmd)
}

func configureApp(cmd *cobra.Command, ac *AppContext) error {
	cmd.PersistentFlags().StringP("config", "c", "", "config file (default is ./sgn.yaml)")
	return nil
}

func preRunApp(cmd *cobra.Command, ac *AppContext) error {
	cfgfile, err := cmd.Flags().GetString("config")
	if err == nil {
		initConfig(cfgfile)
	}
	readConfig()
	for _, run := range ac.PreRuns {
		run()
	}
	return nil
}

func runApp(ac *AppContext) error {
	for _, run := range ac.Runs {
		run()
	}
	return nil
}

func postRunApp(ac *AppContext) error {
	for _, run := range ac.PostRuns {
		run(ac)
	}
	return nil
}
