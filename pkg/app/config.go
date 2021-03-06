package app

import (
	"log"
	"os"
	"os/exec"
	"os/user"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

func initConfig(cfgFile string) {
	addHomePath()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.SetEnvPrefix("SGN")
	viper.AutomaticEnv()
	if cfgFile != "" {
		path, err := filepath.Abs(cfgFile)
		if err == nil {
			viper.SetConfigFile(path)
		}
	}
}

func readConfig() {
	err := viper.ReadInConfig()
	if err == nil {
		log.Println("using config file: ", viper.ConfigFileUsed())
	} else {
		log.Println("read config file error: ", err)
	}
}

func addHomePath() {
	crrDir, err := exec.LookPath(os.Args[0])
	if err == nil {
		crrDir = path.Dir(crrDir)
		viper.AddConfigPath(crrDir)
	}
	u, err := user.Current()
	if err == nil {
		viper.AddConfigPath(path.Join(u.HomeDir, ".sgn"))
	}
}
