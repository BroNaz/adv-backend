package apiserver

// Config - конфигуратор сервиса, заполняется из toml файла
type Config struct {
	BindAddr string `toml:"bint_addr"`
	LogLevel string `toml:"log_level"`
}

// NewConfig - возвращает новый дефолтный конфиг
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
