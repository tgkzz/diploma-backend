package models

import "errors"

var (
	ErrInvalidEmail                error = errors.New("invalid email")
	ErrInvalidPassword             error = errors.New("invalid password")
	ErrEmailAlreadyTaken           error = errors.New("email is already taken")
	ErrIncorrectEmailOrPassword    error = errors.New("incorrect email or password")
	ErrInvalidName                 error = errors.New("firstname and lastname cannot be empty")
	ErrIncorrectUsernameOrPassword error = errors.New("incorrect username or password")
	ErrEmptyness                   error = errors.New("some of the fields may be empty")
)
