package teststore

import (
	"github.com/BroNaz/adv-backend/internal/app/model"
	"github.com/BroNaz/adv-backend/internal/app/store"
)

//Store -  the essence of storage management
type Store struct {
	userRepository *UserRepository
}

// New - return Store
func New() *Store {
	return &Store{}
}

// GetUserRepository - gives you access to manage a table with user data
func (s *Store) GetUserRepository() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}

	return s.userRepository
}
