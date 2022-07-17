package conf

// ApplicationConfig defines the total config
type ApplicationConfig struct {
	// Admin
	AdminConfig *AdminConfig `toml:"admin"`

	// API
	APIConfig *APIConfig `toml:"api"`

	// common
	DatabaseConfig *DatabaseConfig `toml:"database"`
}

type AdminConfig struct {
}

type APIConfig struct {
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
