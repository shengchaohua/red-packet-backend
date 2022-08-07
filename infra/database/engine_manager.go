package database

import (
	"context"

	"xorm.io/xorm"

	"github.com/shengchaohua/red-packet-backend/base/conf"
)

type EngineManager interface {
	GetMasterEngine() *xorm.Engine
	GetSlaveEngine() *xorm.Engine
}

var (
	defaultDBEngineManager EngineManager
)

func InitDBEngineManager(ctx context.Context) {
	defaultDBConfig := conf.GetGlobalAppConfig().DatabaseConfig
	// log

	defaultDBEngineManager = NewDefaultEngineManager(defaultDBConfig)
}

func GetDefaultDBEngineManager() EngineManager {
	return defaultDBEngineManager
}
