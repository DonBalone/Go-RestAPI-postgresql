package main

import (
	"flag"
	"fmt"
	"github.com/DonBalone/Go-RestAPI-postgresql.git/internal/app/apiserver"
	_ "github.com/lib/pq"
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
	/*// открытие бд
	db, err := sql.Open("postgres", "user=username password=password dbname=db sslmode=disable")
	if err != nil {
		log.Fatalf("Error: Unable to connect to database: %v", err)

	}
	defer db.Close()
	// выполнение запросов в бд
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatalf("Error: Unable to execute query: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int64
		var name string
		rows.Scan(&id, &name)
		fmt.Printf("User ID: %d, Name: %s\n", id, name)
	}*/

}
