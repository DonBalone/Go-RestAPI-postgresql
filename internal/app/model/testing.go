package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email:    "user@example.es",
		Password: "password",
	}
}
