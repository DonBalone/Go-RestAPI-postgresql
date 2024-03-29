package main

import (
	"fmt"
	"github.com/DonBalone/Go-RestAPI-postgresql.git/internal/app/apiserver"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
)

func main() {

	// создание нового конфига
	config := apiserver.NewConfig()

	fmt.Println(config)

	// создание нового сервера

	//"postgres://postgres:12345@localhost:5432/restapi_dev?sslmode=disable"
	// если есть ошибка в запуске сервера
	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}

}
