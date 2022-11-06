package config

import (
	"fmt"
	"strings"
)

type Env string

const (
	EnvTest Env = "test"
	EnvLive Env = "live"
)

func mustParseEnv(env string) Env {
	envEnum := Env(strings.ToLower(env))
	switch envEnum {
	case EnvTest, EnvLive:
		return envEnum
	}
	panic(fmt.Errorf("unknown env: %s", env))
}

func (env Env) IsTest() bool {
	return env == EnvTest
}

func (env Env) IsLive() bool {
	return env == EnvLive
}
