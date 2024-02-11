package service

import (
	"fakepayment/internal/repository"
	"fakepayment/internal/service/pay"
)

type Service struct {
	Pay pay.IPayService
}

func NewService(repo repository.Repository, secretKey string) *Service {
	return &Service{
		Pay: pay.NewPayService(repo, secretKey),
	}
}
