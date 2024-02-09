package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (h Handler) Routes() *echo.Echo {
	e := echo.New()

	// CORS settings
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// Rate limiter
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

	courseApi := e.Group("/course")

	courseApi.GET("/", h.getAllPost)
	courseApi.GET("/id/:id", h.getCourseById)
	courseApi.GET("/name/:name", h.getCourseByName)

	courseApi.POST("/create", h.createCourse)

	return e
}
