package app

type AppContext struct {
	CmdBuilder *CommandBuilder
}

func (ac *AppContext) Build() App {
	rootCmd := ac.CmdBuilder.Build()
	return newApp(rootCmd)
}

func (ac *AppContext) RootCmd(use string, description string) {
	ac.CmdBuilder.RootCmd.Use = use
	ac.CmdBuilder.RootCmd.Short = description
}
