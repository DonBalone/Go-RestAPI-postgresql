package main

import (
	"github.com/DonBalone/Go-RestAPI-postgresql.git/internal/app/apiserver"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
)

func main() {

	// создание нового конфига
	config := apiserver.NewConfig()
	// если есть ошибка в запуске сервера
	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}

}

// curl -X POST -H "Content-Type: application/json" -d "{\"email\": \"valid@example.com\", \"password\": \"password\"}" http://localhost:8080/whoami
