package app

type AppContext struct {
	CmdBuilder *CommandBuilder
}

func (ac *AppContext) Build() App {
	rootCmd := ac.CmdBuilder.Build()
	return newApp(rootCmd)
}

func (ac *AppContext) AppInfo(name string, description string) {
	ac.CmdBuilder.RootCmd.Use = name
	ac.CmdBuilder.RootCmd.Short = description
}
