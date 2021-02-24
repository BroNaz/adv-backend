package apiserver

import "github.com/BroNaz/adv-backend/internal/app/store"

// Config - конфигуратор сервиса, заполняется из toml файла
type Config struct {
	BindAddr string `toml:"bint_addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

// NewConfig - возвращает новый дефолтный конфиг
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
