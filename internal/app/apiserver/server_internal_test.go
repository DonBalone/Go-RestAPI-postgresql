package apiserver

import (
	"bytes"
	"encoding/json"
	"github.com/DonBalone/Go-RestAPI-postgresql.git/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_HadleUsersCreate(t *testing.T) {
	s := newServer(teststore.New())
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    "user@example.com",
				"password": "password",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name:         "invalid payload",
			payload:      "userexample.com",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid params",
			payload: map[string]string{
				"email": "userexample.com",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/users", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}

}
