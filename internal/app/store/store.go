package store

import (
	"database/sql"

	_ "github.com/lib/pq" // postgres driver
)

//Store -  the essence of storage management
type Store struct {
	config         *Config
	db             *sql.DB
	userRepository *UserRepository
}

// New - return Store
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

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

// GetUserRepository - gives you access to manage a table with user data
func (s *Store) GetUserRepository() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
