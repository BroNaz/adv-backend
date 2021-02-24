package model

import "testing"

// TestUser - defoult User for testing
func TestUser(t *testing.T) *User {
	return &User{
		Email:    "test@mail.com",
		Password: "qwerty123",
	}
}
