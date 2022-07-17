package mysqlpkg

import "time"

type Config struct {
	// common
	Host     string
	Port     string
	User     string
	Password string
	DBName   string

	// xorm
	DBMaxOpenConns    int
	DBMaxIdleConns    int
	DBConnMaxLifeTime time.Duration
	DBShowSQL         bool
}
