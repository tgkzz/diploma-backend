package handler

import "github.com/gin-gonic/gin"

func (h Handler) Routes() *gin.Engine {
	g := gin.New()

	mainGroup := g.Group("/api")

	//auth service
	authGroup := mainGroup.Group("/auth")
	authGroup.POST("/register", h.register)
	authGroup.POST("/login", h.login)

	//other services

	return g
}
