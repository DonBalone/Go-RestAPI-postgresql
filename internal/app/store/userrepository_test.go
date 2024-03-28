package store_test

import (
	"github.com/DonBalone/Go-RestAPI-postgresql.git/internal/app/model"
	"github.com/DonBalone/Go-RestAPI-postgresql.git/internal/app/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, StorageInfo)

	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email: "user@example.com",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, StorageInfo)

	defer teardown("users")

	email := "user@example.com"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	s.User().Create(&model.User{
		Email: email,
	})
	u, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
