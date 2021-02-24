package model

import (
	"golang.org/x/crypto/bcrypt"
)

//User - representation of the object from the database
type User struct {
	ID                int
	Email             string
	Password          string // virtual atr
	EncryptedPassword string
}

// BeforeCreate ncrypted password
func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}
		u.EncryptedPassword = enc
	}
	return nil
}

func encryptString(str string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
