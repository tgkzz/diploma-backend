package authexpert

import (
	"auth/internal/models"
	"auth/internal/pkg"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func (e ExpertService) CreateExpert(expert models.Expert) error {
	if err := validateExpertData(expert); err != nil {
		return err
	}

	var err error
	expert.Password, err = pkg.HashPassword(expert.Password)
	if err != nil {
		return err
	}

	if err := e.repo.CreateExpert(expert); err != nil {
		return err
	}

	return nil
}

func (e ExpertService) DeleteExpert(email string) error {
	return e.repo.DeleteExpertByEmail(email)
}

func (e ExpertService) GetExpertByEmail(email string) (models.Expert, error) {
	return e.repo.GetExpertByEmail(email)
}

func (e ExpertService) CheckExpertCreds(expert models.Expert) (models.Expert, error) {
	exp, err := e.repo.GetExpertByEmail(expert.Email)
	if err != nil {
		return models.Expert{}, err
	}

	if !pkg.CheckPasswordHash(expert.Password, exp.Password) {
		return models.Expert{}, models.ErrIncorrectUsernameOrPassword
	}

	return exp, nil
}

func (e ExpertService) JwtExpertAuthorization(expert models.Expert) (string, error) {
	claims := models.JwtExpertClaims{
		FirstName: expert.FirstName,
		LastName:  expert.LastName,
		Email:     expert.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			ID:        pkg.RandStringBytesMaskImpr(40),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(e.secretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}
