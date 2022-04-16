package app

import (
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/spf13/viper"
)

func initConfig(cfgFile string) {
	crrDir, err := exec.LookPath(os.Args[0])
	if err != nil {
		log.Fatal(err)
		return
	}
	crrDir = path.Dir(crrDir)
	viper.SetConfigName("nftserver")
	viper.AddConfigPath(crrDir)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	if cfgFile != "" {
		path, err := filepath.Abs(cfgFile)
		if err == nil {
			viper.SetConfigFile(path)
		}
	}
	if err := viper.ReadInConfig(); err == nil {
		log.Println("using config file: ", viper.ConfigFileUsed())
	} else {
		log.Fatal(err)
	}
}
