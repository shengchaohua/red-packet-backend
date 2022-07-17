package database

import (
	"context"

	"xorm.io/xorm"

	"github.com/shengchaohua/red-packet-backend/common/conf"
)

type EngineManager interface {
	GetMasterEngine() *xorm.Engine
	GetSlaveEngine() *xorm.Engine
}

var (
	defaultDBEngineManager EngineManager
)

func InitDBEngineManager(ctx context.Context) {
	defaultDBConfig := conf.GetGlobalConfig().DatabaseConfig
	// log

	defaultDBEngineManager = NewDefaultEngineManager(defaultDBConfig)
}

func GetDefaultDBEngineManager() EngineManager {
	return defaultDBEngineManager
}
