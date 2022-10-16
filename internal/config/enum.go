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

func parseEnv(env string) (Env, error) {
	envEnum := Env(strings.ToLower(env))
	switch envEnum {
	case EnvTest, EnvLive:
		return envEnum, nil
	}
	return "", fmt.Errorf("unknown env: %s", env)
}

func mustParseEnv(env string) Env {
	envEnum, err := parseEnv(env)
	if err != nil {
		panic(err)
	}
	return envEnum
}
