package conf

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

var appConfig *ApplicationConfig

func InitAppConfig(configPath string) error {
	if _, err := toml.DecodeFile(configPath, appConfig); err != nil {
		fmt.Printf("fail to decode config file: %s, error: %s\n", configPath, err.Error())
		return err
	}
	return nil
}

func GetGlobalAppConfig() *ApplicationConfig {
	return appConfig
}
