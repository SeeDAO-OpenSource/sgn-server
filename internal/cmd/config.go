package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type ConfigCmd cobra.Command

func NewConfigCmd() *ConfigCmd {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "配置功能",
		Long:  "配置功能",
		RunE:  func(cmd *cobra.Command, args []string) error { return excuteConfig() },
	}
	return (*ConfigCmd)(cmd)
}

func excuteConfig() error {

	return viper.WriteConfig()
}
