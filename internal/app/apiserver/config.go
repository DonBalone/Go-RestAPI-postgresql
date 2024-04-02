package apiserver

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

// параметры конфига
type Config struct {
	LogLevel    string `yaml:"log_level"`
	HTTPServer  `yaml:"http_server"`
	StorageInfo string `yaml:"storage_info"`
	sessionKey  string `yaml:"session_key"`
}
type HTTPServer struct {
	Address string `yaml:"address"`
}

// создание нового конфига
func NewConfig() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}
	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg

}
