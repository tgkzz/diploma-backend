package models

import "errors"

var (
	ErrInvalidEmail             error = errors.New("invalid email")
	ErrInvalidPassword          error = errors.New("invalid password")
	ErrUsernameAlreadyTaken     error = errors.New("username is already taken")
	ErrEmailAlreadyTaken        error = errors.New("email is already taken")
	ErrIncorrectUsernameOrEmail error = errors.New("incorrect username or email")
)
