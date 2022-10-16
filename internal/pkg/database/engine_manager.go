package database

import (
	"context"

	"go.uber.org/zap"
	"xorm.io/xorm"

	"github.com/shengchaohua/red-packet-backend/internal/config"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/logger"
)

type EngineManager interface {
	GetMasterEngine() *xorm.Engine
	GetSlaveEngine() *xorm.Engine
}

var (
	mainDBEngineManager EngineManager
)

func InitEngineManager(ctx context.Context) {
	mainDBConfig := config.GetGlobalAppConfig().MainDBConfig
	logger.Logger(ctx).Info("[InitEngineManager]initing_main_db", zap.Any("mainDBConfig", mainDBConfig))
	mainDBEngineManager = NewDefaultEngineManager(mainDBConfig)
}

func GetMainDBEngineManager() EngineManager {
	return mainDBEngineManager
}
