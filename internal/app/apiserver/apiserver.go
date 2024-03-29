package apiserver

import (
	"database/sql"
	"github.com/DonBalone/Go-RestAPI-postgresql.git/internal/app/store/sqlstore"
	"net/http"
)

func Start(config *Config) error {
	db, err := newDB(config.StorageInfo)
	if err != nil {
		return err
	}
	defer db.Close()

	store := sqlstore.New(db)
	s := newServer(store)
	return http.ListenAndServe(config.HTTPServer.Address, s)
}

func newDB(storageInfo string) (*sql.DB, error) {
	db, err := sql.Open("postgres", storageInfo)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
