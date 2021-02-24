package store_test

import (
	"testing"

	"github.com/BroNaz/adv-backend/internal/app/model"
	"github.com/BroNaz/adv-backend/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databeseURL)
	defer teardown("users")

	u, err := s.GetUserRepository().Create(&model.User{
		Email:             "user@test.ru",
		EncryptedPassword: "qwerty123",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databeseURL)
	defer teardown("users")

	email := "user@test.ru"
	_, err := s.GetUserRepository().FindByEmail(email)
	assert.Error(t, err)

	s.GetUserRepository().Create(&model.User{
		Email: email,
	})

	u, err := s.GetUserRepository().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
