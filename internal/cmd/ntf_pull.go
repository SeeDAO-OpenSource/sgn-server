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
		RunE:  func(cmd *cobra.Command, args []string) error { return runPull() },
	}
	cmd.PersistentFlags().StringP("address", "a", "", "合约地址")
	cmd.MarkFlagRequired("address")
	viper.BindPFlag("Pull.Address", cmd.PersistentFlags().Lookup("address"))
	return (*NftPullCmd)(cmd)
}

func runPull() error {
	address := viper.GetString("Pull.Address")
	if address == "" {
		return errors.New("参数: address(合约地址)不能为空")
	}
	srv, err := nftv1.BuildNtfServiceV1()
	if err != nil {
		return err
	}
	return srv.PullData(&address, true)
}
