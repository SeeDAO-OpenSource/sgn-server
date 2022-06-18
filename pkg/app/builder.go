package app

import (
	"github.com/SeeDAO-OpenSource/sgn/pkg/utils"
	"github.com/spf13/viper"
)

type AppBuilder struct {
	Context    *AppContext
	CmdBuilder *CommandBuilder
}

func NewAppBuilder() *AppBuilder {
	return &AppBuilder{
		Context:    NewAppContext(),
		CmdBuilder: NewCommandBuilder(),
	}
}

func (ab *AppBuilder) Build() (*App, error) {
	rootCmd, err := ab.CmdBuilder.Build(ab.Context)
	if err != nil {
		return nil, err
	}
	app := newApp(rootCmd)
	return &app, nil
}

func (ab *AppBuilder) Version(version string) {
	ab.Context.Version = version
}

func (ab *AppBuilder) Info(use string, short string, description string) {
	ab.Context.Name = use
	ab.Context.Short = short
	ab.Context.Description = description
}

func (a *AppBuilder) PreRun(action RunFunc) {
	a.Context.PreRun(action)
}

func (a *AppBuilder) Run(action RunFunc) {
	a.Context.Run(action)
}

func (a *AppBuilder) PostRun(action PostRunFunc) {
	a.Context.PostRun(action)
}

func (a *AppBuilder) BindOptions(key string, options interface{}) {
	a.PreRun(func() error {
		utils.BindKey(key, options)
		viper.UnmarshalKey(key, options)
		return nil
	})
}

func (a *AppBuilder) ConfigureServices(action RunFunc) {
	a.Context.PreRun(action)
}
