package cmd

import (
	"errors"

	sgn "github.com/SeeDAO-OpenSource/sgn/internal/apiserver/sgn"
	"github.com/SeeDAO-OpenSource/sgn/pkg/app"
	"github.com/SeeDAO-OpenSource/sgn/pkg/services"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type SgnPullCmd cobra.Command

func NewSgnPullCmd(builder *app.AppBuilder) *SgnPullCmd {
	cmd := &cobra.Command{
		Use:   "pull",
		Short: "拉取sgn信息",
		Long:  "拉取sgn信息",
		RunE:  func(cmd *cobra.Command, args []string) error { return runPull(cmd) },
	}
	cmd.PersistentFlags().StringP("address", "a", "", "合约地址")
	cmd.MarkFlagRequired("address")
	viper.BindPFlag("Pull.Address", cmd.PersistentFlags().Lookup("address"))
	cmd.PersistentFlags().IntP("skip", "s", 0, "跳过数量")
	viper.BindPFlag("Pull.Skip", cmd.PersistentFlags().Lookup("skip"))
	cmd.PersistentFlags().StringArrayP("tokens", "t", []string{}, "token列表")
	return (*SgnPullCmd)(cmd)
}

func runPull(cmd *cobra.Command) error {
	address := viper.GetString("Pull.Address")
	if address == "" {
		return errors.New("参数: address(合约地址)不能为空")
	}
	srv := services.Get[sgn.SgnService]()
	if srv == nil {
		return errors.New("服务: sgn(sgn服务)不能为空")
	}
	skip := viper.GetInt("Pull.Skip")

	tokens, err := cmd.Flags().GetStringArray("tokens")
	if err != nil {
		tokens = []string{}
	}
	return srv.PullData(&address, skip, tokens, true)
}
