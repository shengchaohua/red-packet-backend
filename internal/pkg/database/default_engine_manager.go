package database

import (
	"time"

	"xorm.io/xorm"

	"github.com/shengchaohua/red-packet-backend/internal/config"
	mysqlpkg "github.com/shengchaohua/red-packet-backend/pkg/mysql"
)

type defaultEngineManager struct {
	masterEngine *xorm.Engine
	slaveEngine  *xorm.Engine
}

func (manager *defaultEngineManager) GetMasterEngine() *xorm.Engine {
	return manager.masterEngine
}

func (manager *defaultEngineManager) GetSlaveEngine() *xorm.Engine {
	return manager.slaveEngine
}

func NewDefaultEngineManager(databaseConfig *config.DatabaseConfig) *defaultEngineManager {
	var (
		masterEngine, slaveEngine *xorm.Engine
		err                       error
	)

	for _, dbConfig := range databaseConfig.DBConfigs {
		mysqlConfig := &mysqlpkg.Config{
			Host:              dbConfig.Host,
			Port:              dbConfig.Port,
			User:              dbConfig.User,
			Password:          dbConfig.Password,
			DBName:            dbConfig.DBName,
			DBMaxOpenConns:    dbConfig.DBMaxOpenConns,
			DBMaxIdleConns:    dbConfig.DBMaxIdleConns,
			DBConnMaxLifeTime: time.Duration(dbConfig.DBConnMaxLifeTime) * time.Second,
			DBShowSQL:         dbConfig.DBShowSQL,
		}

		if dbConfig.DBIsMaster {
			if masterEngine != nil {
				panic("master engine has been already inited")
			}

			masterEngine, err = mysqlpkg.NewMySQLEngine(mysqlConfig)
			if err != nil {
				panic(err)
			}
		} else {
			slaveEngine, err = mysqlpkg.NewMySQLEngine(mysqlConfig)
			if err != nil {
				panic(err)
			}
		}
	}

	return &defaultEngineManager{
		masterEngine: masterEngine,
		slaveEngine:  slaveEngine,
	}
}
