package sqlstore_test

import (
	"github.com/DonBalone/Go-RestAPI-postgresql.git/internal/app/model"
	"github.com/DonBalone/Go-RestAPI-postgresql.git/internal/app/store"
	"github.com/DonBalone/Go-RestAPI-postgresql.git/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {

	db, teardown := sqlstore.TestDB(t, StorageInfo)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, StorageInfo)

	defer teardown("users")

	s := sqlstore.New(db)
	email := "user@example.su"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	s.User().Create(u)
	u, err = s.User().FindByEmail(u.Email)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
