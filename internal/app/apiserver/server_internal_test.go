package apiserver

import (
	"github.com/DonBalone/Go-RestAPI-postgresql.git/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_HadleUsersCreate(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/users", nil)
	s := newServer(teststore.New())
	s.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusOK)
}
