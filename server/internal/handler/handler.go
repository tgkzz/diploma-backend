package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"server/internal/service"
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

func (h *Handler) Routes() *echo.Echo {
	e := echo.New()

	// CORS settings
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// Rate limiter
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

	authApi := e.Group("/auth")
	authApi.POST("/register", h.register)
	authApi.POST("/login", h.login)

	authApi.GET("/getUserByEmail", h.getUserByEmail)

	adminApi := e.Group("/admin")
	adminApi.POST("/register", h.registerAdmin)
	adminApi.POST("/login", h.loginAdmin)

	expertApi := e.Group("/expert")
	expertApi.POST("/register", h.registerExpert)
	expertApi.POST("/login", h.loginExpert)
	expertApi.GET("/getAllExperts", h.getAllExperts)

	return e
}
