package apiserver

type Config struct {
	BindAddr string `toml:"bint_addr"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
	}
}