package pay

import (
	"fakepayment/internal/model"
	"fakepayment/internal/repository/pay"
)

type PayService struct {
	repo      pay.IPayRepo
	secretKey string
}

type IPayService interface {
	BuyCourse(input model.ClientInput) error
}

func NewPayService(repo pay.IPayRepo, secretKey string) *PayService {
	return &PayService{
		repo:      repo,
		secretKey: secretKey,
	}
}
