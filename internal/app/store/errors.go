package store

import "errors"

var (
	ErrRecordNotFound = errors.New("not found user with such an email")
)
