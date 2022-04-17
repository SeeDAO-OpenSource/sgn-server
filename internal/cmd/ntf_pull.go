package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	nftv1 "github.com/waite-lee/nftserver/internal/apiserver/nft/v1"
)

type NftPullCmd cobra.Command

func NewNftPullCmd() *NftPullCmd {

	cmd := &cobra.Command{
		Use:   "pull",
		Short: "拉取nft信息",
		Long:  "拉取nft信息",
		RunE:  func(cmd *cobra.Command, args []string) error { return runPull(cmd) },
	}
	cmd.PersistentFlags().StringP("address", "a", "", "合约地址")
	cmd.MarkFlagRequired("address")
	viper.BindPFlag("Pull.Address", cmd.PersistentFlags().Lookup("address"))
	cmd.PersistentFlags().IntP("skip", "s", 0, "跳过数量")
	viper.BindPFlag("Pull.Skip", cmd.PersistentFlags().Lookup("skip"))
	cmd.PersistentFlags().StringArrayP("tokens", "t", []string{}, "token列表")
	return (*NftPullCmd)(cmd)
}

func runPull(cmd *cobra.Command) error {
	address := viper.GetString("Pull.Address")
	if address == "" {
		return errors.New("参数: address(合约地址)不能为空")
	}
	srv, err := nftv1.BuildNftServiceV1()
	if err != nil {
		return err
	}
	skip := viper.GetInt("Pull.Skip")

	tokens, err := cmd.Flags().GetStringArray("tokens")
	if err != nil {
		tokens = []string{}
	}
	return srv.PullData(&address, skip, tokens, true)
}
