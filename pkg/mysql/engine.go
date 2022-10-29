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
		Addr:                 host,
		User:                 config.User,
		Passwd:               config.Password,
		DBName:               config.DBName,
		AllowNativePasswords: config.AllowNativePasswords,
	}

	engine, err := xorm.NewEngine("mysql", mysqlConfig.FormatDSN())
	if err != nil {
		return nil, err
	}

	engine.SetMaxOpenConns(config.MaxOpenConns)
	engine.SetMaxIdleConns(config.MaxIdleConns)
	engine.SetConnMaxLifetime(config.ConnMaxLifeTime)
	engine.ShowSQL(config.ShowSQL)

	return engine, nil
}
