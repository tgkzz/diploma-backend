package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) Routes() *echo.Echo {
	e := echo.New()

	authApi := e.Group("/auth")
	authApi.POST("/register", h.register)
	authApi.POST("/login", h.login)

	adminApi := e.Group("/admin")
	adminApi.POST("/register", h.registerAdmin)
	adminApi.POST("/login", h.loginAdmin)

	expertApi := e.Group("/expert")
	expertApi.POST("/register", h.registerExpert)
	expertApi.POST("/login", h.loginExpert)

	return e
}
