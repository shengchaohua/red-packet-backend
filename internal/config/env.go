package config

import (
	"fmt"
	"strings"
)

type Env string

const (
	EnvDev  Env = "dev"
	EnvLive Env = "live"
)

func mustParseEnv(env string) Env {
	envEnum := Env(strings.ToLower(env))
	switch envEnum {
	case EnvDev, EnvLive:
		return envEnum
	}
	panic(fmt.Errorf("unknown env: %s", env))
}

func (env Env) IsDev() bool {
	return env == EnvDev
}

func (env Env) IsLive() bool {
	return env == EnvLive
}
