package store

// Config - the main configuration for working with the database
type Config struct {
	DatabaseURL string `toml:"database_url"`
}

// NewConfig - returns the config
func NewConfig() *Config {
	return &Config{
		//DatabaseURL: "host=localhost dbname=restapi_dev sslmode=disable",
	}
}
