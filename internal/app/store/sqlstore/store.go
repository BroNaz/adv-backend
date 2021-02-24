package sqlstore

import (
	"database/sql"

	"github.com/BroNaz/adv-backend/internal/app/store"
	_ "github.com/lib/pq" // postgres driver
)

//Store -  the essence of storage management
type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

// New - return Store
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// GetUserRepository - gives you access to manage a table with user data
func (s *Store) GetUserRepository() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
