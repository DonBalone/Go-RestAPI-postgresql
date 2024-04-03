package store

import "github.com/DonBalone/Go-RestAPI-postgresql.git/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
	FindById(int) (*model.User, error)
}
