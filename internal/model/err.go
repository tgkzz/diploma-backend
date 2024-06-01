package model

import "errors"

var (
	ErrInvalidEmail                error = errors.New("invalid email")
	ErrInvalidPassword             error = errors.New("invalid password")
	ErrEmailAlreadyTaken           error = errors.New("email is already taken")
	ErrIncorrectEmailOrPassword    error = errors.New("incorrect email or password")
	ErrInvalidName                 error = errors.New("firstname and lastname cannot be empty")
	ErrIncorrectUsernameOrPassword error = errors.New("incorrect username or password")
	ErrEmptiness                   error = errors.New("some of the fields may be empty")
	ErrIncorrectCode               error = errors.New("incorrect code")
	ErrNotFound                    error = errors.New("not found")
	ErrCourseAlreadyPurchased      error = errors.New("course already purchased")
	ErrEmailIsAlreadyUser          error = errors.New("email is already used")
	ErrNoMeeting                   error = errors.New("no room meeting")
	ErrTimeInPast                  error = errors.New("time cannot be in past")
	ErrMeetingAlreadyBooked        error = errors.New("meeting already token")
)
