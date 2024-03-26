package apiserver

// параметры конфига
type Config struct {
	BindAddr string `toml:"bind_addr"`
}

// создание нового конфига
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
	}
}
