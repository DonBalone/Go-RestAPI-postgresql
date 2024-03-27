package apiserver

import "github.com/DonBalone/Go-RestAPI-postgresql.git/internal/app/store"

// параметры конфига
type Config struct {
	LogLevel   string `yaml:"log_level"`
	HTTPServer `yaml:"http_server"`
	Store      *store.Config
}
type HTTPServer struct {
	Address string `yaml:"address"`
}

// создание нового конфига
func NewConfig() *Config {
	return &Config{
		LogLevel: "debug",
		Store:    store.NewConfig(),
		HTTPServer: HTTPServer{
			Address: "localhost:8080",
		},
	}
}
