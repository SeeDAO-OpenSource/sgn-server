package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/waite-lee/sgn/internal/apiserver"
)

type ApiServerCmd cobra.Command

func NewApiServerCmd(server apiserver.ApiServer, options *apiserver.ApiServerOptions) *ApiServerCmd {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "启动文档服务",
		Long:  "启动一个文档服务",
		RunE:  func(cmd *cobra.Command, args []string) error { return runAction(&server, options) },
	}
	cmd.PersistentFlags().IntP("port", "p", options.Port, "端口号")
	viper.BindPFlag("apiserver.port", cmd.PersistentFlags().Lookup("port"))
	return (*ApiServerCmd)(cmd)
}

func runAction(server *apiserver.ApiServer, options *apiserver.ApiServerOptions) error {
	viper.UnmarshalKey("ApiServer", &options)
	return server.Run()
}
