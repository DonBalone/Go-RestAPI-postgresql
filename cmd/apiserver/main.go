package main

import (
	"flag"
	"fmt"
	"github.com/DonBalone/Go-RestAPI-postgresql.git/internal/app/apiserver"

	//_ "github.com/lib/pq"
	"io/ioutil"
	"log"
)

func main() {
	// разделяем флаги и получаем все значения для конфига
	flag.Parse()
	configPath := "configs/apiserver.yaml"
	// создание нового конфига
	config := apiserver.NewConfig()

	fmt.Println(config)
	_, err := ioutil.ReadFile(configPath)
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
