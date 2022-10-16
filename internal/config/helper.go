package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

var appConfig = new(ApplicationConfig)

func InitAppConfig(configPath string) {
	log.Println("Initing config with file:", configPath)
	if _, err := toml.DecodeFile(configPath, appConfig); err != nil {
		log.Fatalf("fail to decode config file: %s, error: %s", configPath, err.Error())
	}

	mustParseEnv(appConfig.AdminConfig.Env)
	mustParseEnv(appConfig.APIConfig.Env)
}

func GetGlobalAppConfig() *ApplicationConfig {
	return appConfig
}
