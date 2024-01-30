package authexpert

import (
	"auth/internal/models"
	"auth/internal/repository/authexpert"
)

type ExpertService struct {
	repo authexpert.IExpertRepo
}

type IExpertService interface {
	CreateExpert(expert models.Expert) error
	DeleteExpert(email string) error
	GetExpertByEmail(email string) (models.Expert, error)
	CheckExpertCreds(expert models.Expert) (models.Expert, error)
	JwtExpertAuthorization(expert models.Expert) (string, error)
}

func NewExpertService(repo authexpert.IExpertRepo) *ExpertService {
	return &ExpertService{
		repo: repo,
	}
}
