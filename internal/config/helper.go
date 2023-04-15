package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

var appConfig = new(ApplicationConfig)

func InitConfig(configPath string) {
	log.Println("[Info]Init config with file:", configPath)
	if _, err := toml.DecodeFile(configPath, appConfig); err != nil {
		log.Fatalf("fail to decode config file: %s, error: %s", configPath, err.Error())
	}

	envEnum := mustParseEnv(appConfig.ServerConfig.Env)
	appConfig.ServerConfig.EnvEnum = envEnum

	roleEnum := mustParseRole(appConfig.ServerConfig.Role)
	appConfig.ServerConfig.RoleEnum = roleEnum
}

func GetGlobalConfig() *ApplicationConfig {
	return appConfig
}
