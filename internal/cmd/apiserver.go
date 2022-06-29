package cmd

import (
	"github.com/SeeDAO-OpenSource/sgn/internal/apiserver"
	"github.com/SeeDAO-OpenSource/sgn/pkg/app"
	"github.com/SeeDAO-OpenSource/sgn/pkg/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type ApiServerCmd cobra.Command

func NewApiServerCmd(builder *app.AppBuilder) *ApiServerCmd {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "SGN Api Server",
		Long:  "启动SGN服务",
	}
	serverBuiler := server.NewServerBuilder(builder)
	apiserver.AddApiServer(serverBuiler)
	options := serverBuiler.Options
	cmd.PersistentFlags().IntVarP(&options.Port, "port", "p", options.Port, "端口号")
	viper.BindPFlag("apiserver.port", cmd.PersistentFlags().Lookup("port"))
	cmd.RunE = func(cmd *cobra.Command, args []string) error { return runAction(serverBuiler) }
	return (*ApiServerCmd)(cmd)
}

func runAction(builder *server.ServerBuiler) error {
	s, err := builder.Build()
	if err != nil {
		return err
	}
	return s.Run()
}
