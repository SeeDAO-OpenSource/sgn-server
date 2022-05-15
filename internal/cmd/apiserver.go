package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/waite-lee/sgn/internal/apiserver"
	"github.com/waite-lee/sgn/pkg/app"
	"github.com/waite-lee/sgn/pkg/server"
)

type ApiServerCmd cobra.Command

func NewApiServerCmd(builder *app.AppBuilder) *ApiServerCmd {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "SGN Api Server",
		Long:  "启动SGN服务",
	}
	options := &server.ServerOptions{}
	serverBuiler := server.AddServer(builder, options)
	apiserver.AddApiServer(serverBuiler)
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
