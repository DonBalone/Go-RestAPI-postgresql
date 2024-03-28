package store

import (
	"database/sql"
	_ "github.com/lib/pq"
)

// это сущность хранилища, которая
// скрывает все детали реализации
type Store struct {
	config *Config
	db     *sql.DB
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// нужен для инициализации
// хранилища, при подключении к бд
func (s *Store) Open() error {

	db, err := sql.Open("postgres", s.config.StorageInfo)

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}
	// записываем нашу бд в поле db
	s.db = db

	return nil
}

// когда веб сервер заканчивает работу
// чтобы мы могли отключиться от бд
func (s *Store) Close() {

}

func (s *Store) Config() {
	s.db.Close()
}
