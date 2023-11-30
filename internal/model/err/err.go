package err

import "errors"

type Err struct {
	Text string
	Code int
}

var (
	ErrInvalidData       error = errors.New("invalid email or password")
	ErrIncorrectPassword error = errors.New("incorrect password")
	ErrInvalidToken      error = errors.New("invalid token")
)
