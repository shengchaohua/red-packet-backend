package config

// ApplicationConfig defines the total config
type ApplicationConfig struct {
	MainDBConfig *DatabaseConfig `toml:"main_db"`
	ServerConfig *ServerConfig   `toml:"server"`
}

type ServerConfig struct {
	Env     string `toml:"env"`
	Port    string `toml:"port"`
	LogFile string `toml:"log_file"`
	Role    string `toml:"role"`

	EnvEnum  Env
	RoleEnum Role
}

func (serverConfig *ServerConfig) IsDevEnv() bool {
	return serverConfig.EnvEnum.IsDev()
}

func (serverConfig *ServerConfig) IsLiveEnv() bool {
	return serverConfig.EnvEnum.IsLive()
}

func (serverConfig *ServerConfig) IsAdmin() bool {
	return serverConfig.RoleEnum.IsAdmin()
}

func (serverConfig *ServerConfig) IsApi() bool {
	return serverConfig.RoleEnum.IsApi()
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
