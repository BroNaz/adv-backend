package store

import "github.com/BroNaz/adv-backend/internal/app/model"

// UserRepository - entity for managing a table with users
type UserRepository struct {
	store *Store
}

// Create accepts a standard structure with the fields email, encrypted_password,
// creates an object in the database and returns a fully prepared user view
func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := u.BeforeCreate(); err != nil {
		return nil, err
	}

	err := r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID)
	if err != nil {
		return nil, err
	}

	return u, nil
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
		return nil, err
	}
	return users, nil
}
