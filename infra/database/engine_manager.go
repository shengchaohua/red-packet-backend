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
	mainDBEngineManager EngineManager
)

func InitDBEngineManager(ctx context.Context) {
	mainDBConfig := conf.GetGlobalAppConfig().MainDBConfig
	// TODO log
	mainDBEngineManager = NewDefaultEngineManager(mainDBConfig)
}

func GetMainBEngineManager() EngineManager {
	return mainDBEngineManager
}
