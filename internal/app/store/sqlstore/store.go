package sqlstore

import (
	"database/sql"
	"github.com/DonBalone/Go-RestAPI-postgresql.git/internal/app/store"
	_ "github.com/lib/pq"
)

// это сущность хранилища, которая
// скрывает все детали реализации
type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}

}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
