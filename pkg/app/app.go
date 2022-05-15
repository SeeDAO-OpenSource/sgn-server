package app

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/waite-lee/sgn/pkg/options"
)

type App struct {
	rootCmd       *cobra.Command
	optionsInitor *options.OptionsInitor
}

func newApp(root *cobra.Command) App {
	return App{
		rootCmd: root,
	}
}

func (app App) Run() {
	if err := app.rootCmd.Execute(); err != nil {
		log.Fatal("程序出错: " + err.Error())
		os.Exit(1)
	}
}
