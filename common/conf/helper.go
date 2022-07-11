package conf

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

func InitAppConfig(configPath string) error {
	if _, err := toml.DecodeFile(configPath, appConfig); err != nil {
		fmt.Printf("fail to decode config file: %s, error: %s", configPath, err.Error())
		return err
	}
	return nil
}

func GetGlobalConfig() *ApplicationConfig {
	return appConfig
}
