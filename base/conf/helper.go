package conf

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

var appConfig = new(ApplicationConfig)

func InitAppConfig(configPath string) {
	if _, err := toml.DecodeFile(configPath, appConfig); err != nil {
		fmt.Printf("fail to decode config file: %s, error: %s\n", configPath, err.Error())
		panic(err)
	}
}

func GetGlobalAppConfig() *ApplicationConfig {
	return appConfig
}
