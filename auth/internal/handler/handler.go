package handler

import (
	"auth/internal/service"
	"log"
)

type Handler struct {
	service     *service.Service
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

func NewHandler(service *service.Service, info *log.Logger, err *log.Logger) *Handler {
	return &Handler{
		service:     service,
		infoLogger:  info,
		errorLogger: err,
	}
}
