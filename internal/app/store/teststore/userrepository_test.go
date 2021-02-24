package teststore_test

import (
	"testing"

	"github.com/BroNaz/adv-backend/internal/app/model"
	"github.com/BroNaz/adv-backend/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {

	s := teststore.New()
	u := model.TestUser(t)
	err := s.GetUserRepository().Create(u)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {

	s := teststore.New()
	email := "user@test.ru"
	_, err := s.GetUserRepository().FindByEmail(email)
	assert.Error(t, err)

	user := model.TestUser(t)
	user.Email = email
	s.GetUserRepository().Create(user)

	u, err := s.GetUserRepository().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
