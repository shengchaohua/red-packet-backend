package usergrouprelationdm

import (
	"github.com/shengchaohua/red-packet-backend/internal/config"
)

const (
	shardingNumberTestEnv = 2
	shardingNumberLiveEnv = 10
)

func getShardingNumberByEnv(env config.Env) uint64 {
	if env.IsLive() {
		return shardingNumberLiveEnv
	}
	if env.IsDev() {
		return shardingNumberTestEnv
	}
	return shardingNumberLiveEnv // use live env as default
}
