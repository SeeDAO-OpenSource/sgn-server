package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

type TestCmd cobra.Command

func NewTestCmd() *TestCmd {
	cmd := &cobra.Command{
		Use:   "test",
		Short: "test",
		Long:  "test",
		RunE:  func(cmd *cobra.Command, args []string) error { return excute() },
	}
	return (*TestCmd)(cmd)
}

func excute() error {
	log.Println("测试")
	return nil
}
