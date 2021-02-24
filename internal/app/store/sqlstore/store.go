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

/*
// Open - opens and pings the database
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db
	return nil
}

// Close - close db
func (s *Store) Close() {
	s.db.Close()
}
*/

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
