package authexpert

import (
	"github.com/golang-jwt/jwt/v5"
	"server/internal/model"
	"server/internal/pkg"
	"server/internal/repository/authexpert"
	"time"
)

type ExpertService struct {
	repo      authexpert.IExpertRepo
	secretKey string
}

type IExpertService interface {
	CreateExpert(expert model.Expert) error
	DeleteExpert(email string) error
	GetExpertByEmail(email string) (model.Expert, error)
	GetAllExperts() ([]model.Expert, error)
	CheckExpertCreds(expert model.Expert) (model.Expert, error)
	JwtExpertAuthorization(expert model.Expert) (string, error)
}

func NewExpertService(repo authexpert.IExpertRepo, secretKey string) *ExpertService {
	return &ExpertService{
		repo:      repo,
		secretKey: secretKey,
	}
}

func validateExpertData(expert model.Expert) error {
	if !pkg.IsValid(expert) {
		return model.ErrEmptiness
	}

	if !pkg.IsEmailValid(expert.Email) {
		return model.ErrInvalidEmail
	}

	if !pkg.IsNameValid(expert.FirstName, expert.LastName) {
		return model.ErrInvalidName
	}

	if !pkg.IsPasswordStrong(expert.Password) {
		return model.ErrInvalidPassword
	}

	return nil
}

func (e ExpertService) CreateExpert(expert model.Expert) error {
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

func (e ExpertService) GetExpertByEmail(email string) (model.Expert, error) {
	return e.repo.GetExpertByEmail(email)
}

func (e ExpertService) GetAllExperts() ([]model.Expert, error) {
	return e.repo.GetAllExpert()
}

func (e ExpertService) CheckExpertCreds(expert model.Expert) (model.Expert, error) {
	exp, err := e.repo.GetExpertByEmail(expert.Email)
	if err != nil {
		return model.Expert{}, err
	}

	if !pkg.CheckPasswordHash(expert.Password, exp.Password) {
		return model.Expert{}, model.ErrIncorrectUsernameOrPassword
	}

	return exp, nil
}

func (e ExpertService) JwtExpertAuthorization(expert model.Expert) (string, error) {
	claims := model.JwtExpertClaims{
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
