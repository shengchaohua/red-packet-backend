package config

// ApplicationConfig defines the total config
type ApplicationConfig struct {
	MainDBConfig *DatabaseConfig `toml:"main_db"`
	ServerConfig *ServerConfig   `toml:"server"`
}

type ServerConfig struct {
	Addr string `toml:"addr"`
	Env  string `toml:"env"`
	Log  string `toml:"log"`
	Role string `toml:"role"`
}

func (serverConfig *ServerConfig) GetEnv() Env {
	return mustParseEnv(serverConfig.Env)
}

func (serverConfig *ServerConfig) IsTestEnv() bool {
	return serverConfig.GetEnv().IsTest()
}

func (serverConfig *ServerConfig) IsLiveEnv() bool {
	return mustParseEnv(serverConfig.Env).IsLive()
}

func (serverConfig *ServerConfig) GetRole() Role {
	return mustParseRole(serverConfig.Role)
}

func (serverConfig *ServerConfig) IsAdmin() bool {
	return serverConfig.GetRole().IsAdmin()
}

func (serverConfig *ServerConfig) IsAPI() bool {
	return serverConfig.GetRole().IsAPI()
}

type DatabaseConfig struct {
	DBConfigs []*DBConfig `toml:"db"`
}

type DBConfig struct {
	Host                 string `toml:"host"`
	Port                 string `toml:"port"`
	User                 string `toml:"user"`
	Password             string `toml:"password"`
	AllowNativePasswords bool   `toml:"allow_native_passwords"`
	DBName               string `toml:"db_name"`
	DBMaxOpenConns       int    `toml:"db_max_open_conns"`
	DBMaxIdleConns       int    `toml:"db_max_idle_conns"`
	DBConnMaxLifeTime    int    `toml:"db_max_life_time"` // seconds
	DBShowSQL            bool   `toml:"db_show_sql"`
	DBIsMaster           bool   `toml:"db_is_master"`
}
