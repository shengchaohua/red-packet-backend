package config

// ApplicationConfig defines the total config
type ApplicationConfig struct {
	// admin
	AdminConfig *ServerConfig `toml:"admin"`

	// api
	APIConfig *ServerConfig `toml:"api"`

	// common
	MainDBConfig *DatabaseConfig `toml:"main_db"`
}

type ServerConfig struct {
	Addr string `toml:"addr"`
	Env  string `toml:"env"`
	Log  string `toml:"log"`
}

func (serverConfig *ServerConfig) IsTestEnv() bool {
	return mustParseEnv(serverConfig.Env) == EnvTest
}

func (serverConfig *ServerConfig) IsLiveEnv() bool {
	return mustParseEnv(serverConfig.Env) == EnvLive
}

type DatabaseConfig struct {
	DBConfigs []*DBConfig `toml:"db"`
}

type DBConfig struct {
	Host              string `toml:"host"`
	Port              string `toml:"port"`
	User              string `toml:"user"`
	Password          string `toml:"password"`
	DBName            string `toml:"db_name"`
	DBMaxOpenConns    int    `toml:"db_max_open_conns"`
	DBMaxIdleConns    int    `toml:"db_max_idle_conns"`
	DBConnMaxLifeTime int    `toml:"db_max_life_time"` // seconds
	DBShowSQL         bool   `toml:"db_show_sql"`
	DBIsMaster        bool   `toml:"db_is_master"`
}
