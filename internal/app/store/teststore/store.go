package teststore

import (
	"github.com/DonBalone/Go-RestAPI-postgresql.git/internal/app/model"
	_ "github.com/lib/pq"
)

// это сущность хранилища, которая
// скрывает все детали реализации
type Store struct {
	userRepository *UserRepository
}

func New() *Store {
}

func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}

	return s.userRepository
}
