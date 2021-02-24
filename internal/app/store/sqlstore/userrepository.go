package sqlstore

import (
	"database/sql"

	"github.com/BroNaz/adv-backend/internal/app/model"
	"github.com/BroNaz/adv-backend/internal/app/store"
)

// UserRepository - entity for managing a table with users
type UserRepository struct {
	store *Store
}

// Create accepts a standard structure with the fields email, encrypted_password,
// creates an object in the database and returns a fully prepared user view
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID)
}

// FindByEmail search for the user by email
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	users := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email = $1",
		email,
	).Scan(
		&users.ID,
		&users.Email,
		&users.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return users, nil
}
