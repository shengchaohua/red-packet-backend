package usergroupmappingdm

import (
	"github.com/shengchaohua/red-packet-backend/internal/config"
)

const (
	shardingNumberTestEnv = 2
	shardingNumberLiveEnv = 1000
)

func getShardingNumberByEnv(env config.Env) uint64 {
	switch env {
	case config.EnvLive:
		return shardingNumberLiveEnv
	case config.EnvTest:
		return shardingNumberTestEnv
	}
	return shardingNumberLiveEnv // use live env as default
}
