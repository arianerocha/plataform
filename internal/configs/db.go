package configs

// DbConfig struct
type DbConfig struct{}

// Db config
func Db() *DbConfig {
	return &DbConfig{}
}

// Adapter Db
func (c *DbConfig) Adapter() string {
	return Sys().Env("DB_ADAPTER", "sqlite3")
}

// URL for the Db connection
func (c *DbConfig) URL() string {
	return Sys().Env("DB_URL", "/tmp/klab_plataform.db")
}
