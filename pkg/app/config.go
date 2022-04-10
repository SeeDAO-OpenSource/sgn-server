package app

import (
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/spf13/viper"
)

var cfgFile = "$HOME/.ntfserver.yaml"

func init() {
	initConfig()
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		crrDir, err := exec.LookPath(os.Args[0])
		if err != nil {
			log.Fatal(err)
			return
		}
		crrDir = path.Dir(crrDir)
		viper.AddConfigPath(crrDir)
		viper.SetConfigType("yaml")
		viper.SetConfigName("conf")
	}
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		log.Println("using config file: ", viper.ConfigFileUsed())
	}
}
