package auth

import (
	erroring "diploma/internal/model/err"
	"diploma/internal/model/user"
	"diploma/internal/pkg"
)

func (s *AuthService) CreateUser(user user.User) error {
	var err error

	if !pkg.DataValidation(user) {
		return erroring.ErrInvalidData
	}

	user.Password, err = pkg.HashPassword(user.Password)
	if err != nil {
		return err
	}

	return s.repo.CreateUser(user)
}

func (s *AuthService) CheckUserCreds(creds user.User) error {
	user, err := s.repo.GetUserByLogin(creds.Login)
	if err != nil {
		return err
	}

	if !pkg.CheckPasswordHash(creds.Password, user.Password) {
		return erroring.ErrIncorrectPassword
	}
	return nil
}

func (s *AuthService) CreateToken(login string) (string, error) {
	token, err := pkg.CreateToken(login)
	if err != nil {
		return "", err
	}

	return token, nil
}
