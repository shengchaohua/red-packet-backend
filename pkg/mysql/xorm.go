package mysqlpkg

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func NewMySQLEngine(config *Config) (*xorm.Engine, error) {
	host := config.Host
	if len(config.Port) != 0 {
		host = fmt.Sprintf("%s:%s", config.Host, config.Port)
	}

	mysqlConfig := &mysql.Config{
		Addr:   host,
		User:   config.User,
		Passwd: config.Password,
		DBName: config.DBName,
	}

	engine, err := xorm.NewEngine("mysql", mysqlConfig.FormatDSN())
	if err != nil {
		return nil, err
	}

	engine.SetMaxOpenConns(config.DBMaxOpenConns)
	engine.SetMaxIdleConns(config.DBMaxIdleConns)
	engine.SetConnMaxLifetime(config.DBConnMaxLifeTime)
	engine.ShowSQL(config.DBShowSQL)

	return engine, nil
}
