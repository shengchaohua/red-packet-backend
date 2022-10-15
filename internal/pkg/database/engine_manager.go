package database

import (
	"context"

	"xorm.io/xorm"

	"github.com/shengchaohua/red-packet-backend/internal/config"
)

type EngineManager interface {
	GetMasterEngine() *xorm.Engine
	GetSlaveEngine() *xorm.Engine
}

var (
	mainDBEngineManager EngineManager
)

func InitDBEngineManager(ctx context.Context) {
	mainDBConfig := config.GetGlobalAppConfig().MainDBConfig
	// TODO log
	mainDBEngineManager = NewDefaultEngineManager(mainDBConfig)
}

func GetMainBEngineManager() EngineManager {
	return mainDBEngineManager
}
