package config

import (
	"github.com/spf13/viper"
	"os"
	"fmt"
)

func dirExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fi.IsDir()
}

func InitConfig(name string) error{
	viper.SetConfigName(name)
	var altPath = os.Getenv("MINILEDGER_CFG_PATH")
	if altPath != "" {
		//if the env param is given, it should be considered first
		if !dirExists(altPath) {
			return fmt.Errorf("MINILEDGER_CFG_PATH %s does not exist", altPath)
		}
		viper.AddConfigPath(altPath)
	} else {
		viper.AddConfigPath(".")
		viper.AddConfigPath(".")
		if dirExists(OfficialPath) {
			AddConfigPath(v, OfficialPath)
		}
	}

	// Now set the configuration file.
	if v != nil {
		v.SetConfigName(configName)
	} else {
		viper.SetConfigName(configName)
	}

	return nil
}
