package main

import (
	"flag"
	"github.com/BurntSushi/toml"

	"github.com/DonBalone/Go-RestAPI-postgresql.git/internal/app/apiserver"
	"log"
)

var (
	configPath string
)

func main() {
	// разделяем флаги и получаем все значения для конфига
	flag.Parse()
	// создание нового конфига
	config := apiserver.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	// создание нового сервера
	s := apiserver.New(config)
	// если есть ошибка в запуске сервера
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}

// создание флага, который получает путь к конфигу
func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}
