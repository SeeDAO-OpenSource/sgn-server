package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

type TestCmd cobra.Command

func NewTestCmd() *TestCmd {
	cmd := &cobra.Command{
		Use:   "test",
		Short: "测试功能",
		Long:  "测试功能",
		RunE:  func(cmd *cobra.Command, args []string) error { return excute() },
	}
	return (*TestCmd)(cmd)
}

func excute() error {
	log.Println("测试")
	//client, err := ethclient.Dial("https://ropsten.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161")
	return nil
}
