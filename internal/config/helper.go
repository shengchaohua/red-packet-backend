package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

var appConfig = new(ApplicationConfig)

func InitAppConfig(configPath string) {
	if _, err := toml.DecodeFile(configPath, appConfig); err != nil {
		log.Fatalf("fail to decode config file: %s, error: %s", configPath, err.Error())
	}
}

func GetGlobalAppConfig() *ApplicationConfig {
	return appConfig
}
