package cmd

import (
	"log"
	"os"

	"github.com/SeeDAO-OpenSource/sgn/pkg/db/mongodb"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type TestCmd cobra.Command

func NewTestCmd() *TestCmd {
	cmd := &cobra.Command{
		Use:   "test",
		Short: "test",
		Long:  "test",
		RunE:  func(cmd *cobra.Command, args []string) error { return test() },
	}
	return (*TestCmd)(cmd)
}

func test() error {
	log.Println("测试")
	v := os.Getenv("SGN_PULL_ADDRESS")
	log.Println(v)
	gs := viper.GetString("Mongo.URL")
	log.Println(gs)
	options := mongodb.MongoOptions{}
	viper.UnmarshalKey("Mongo", &options)
	log.Println(options.URL)
	return nil
}
