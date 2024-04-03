package apiserver

import (
	"database/sql"
	"github.com/DonBalone/Go-RestAPI-postgresql.git/internal/app/store/sqlstore"
	"github.com/gorilla/sessions"
	"net/http"
)

func Start(config *Config) error {
	db, err := newDB(config.StorageInfo)
	if err != nil {
		return err
	}
	defer db.Close()
	config.sessionKey = "81099f8b6b5e4b0e8c486b5e4b0e8c48"
	store := sqlstore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(config.sessionKey))
	s := newServer(store, sessionStore)
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
