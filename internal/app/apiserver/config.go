package apiserver

// Config - конфигуратор сервиса, заполняется из toml файла
type Config struct {
	BindAddr    string `toml:"bint_addr"`
	LogLevel    string `toml:"log_level"`
	DatabaseURL string `toml:database_url`
}

// NewConfig - возвращает новый дефолтный конфиг
func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8080",
		LogLevel:    "debug",
		DatabaseURL: "host=localhost dbname=restapi_dev sslmode=disable user=postgres password=12",
	}
}
