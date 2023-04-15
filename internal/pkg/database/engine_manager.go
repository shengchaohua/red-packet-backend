package database

import (
	"context"
	"github.com/shengchaohua/red-packet-backend/internal/config"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/logger"
	"go.uber.org/zap"
	"xorm.io/xorm"
)

type EngineManager interface {
	GetMasterEngine() *xorm.Engine
	GetSlaveEngine() *xorm.Engine
}

var (
	mainDBEngineManager EngineManager
)

func InitEngineManager(ctx context.Context) {
	mainDBConfig := config.GetGlobalConfig().MainDBConfig
	logger.Logger(ctx).Info("Init main DB", zap.Any("mainDBConfig", mainDBConfig))
	mainDBEngineManager = NewDefaultEngineManager(mainDBConfig)
}

func GetMainDBEngineManager() EngineManager {
	if mainDBEngineManager == nil {
		panic("mainDBEngineManager is nil")
	}
	return mainDBEngineManager
}
