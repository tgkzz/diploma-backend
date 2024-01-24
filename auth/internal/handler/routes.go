package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) Routes() *echo.Echo {
	e := echo.New()

	authApi := e.Group("/auth")

	authApi.POST("/register", h.register)
	authApi.POST("/login", h.login)

	return e
}
