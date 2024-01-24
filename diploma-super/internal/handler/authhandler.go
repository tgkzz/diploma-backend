package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"super/internal/models"
	"super/internal/models/auth"
)

func (h *Handler) register(c *gin.Context) {
	user := models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		h.errLogger.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result, err := h.service.Auth.CreateNewUser(user)
	if err != nil {
		h.errLogger.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result["status"] == "fail" {
		switch result["message"] {
		case auth.ErrInvalidEmail.Error(), auth.ErrInvalidPassword.Error():
			c.JSON(http.StatusBadRequest, gin.H{"message": result["message"]})
		case auth.ErrUsernameAlreadyTaken.Error():
			c.JSON(http.StatusConflict, gin.H{"message": result["message"]})
		case auth.ErrEmailAlreadyTaken.Error():
			c.JSON(http.StatusConflict, gin.H{"message": result["message"]})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": "unknown error"})
		}
		return
	}

	h.infoLogger.Print("successfully created new user")
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "successfully created new user"})
}

func (h *Handler) login(c *gin.Context) {
	user := models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		h.errLogger.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result, err := h.service.Auth.Login(user)
	if err != nil {
		h.errLogger.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result["status"] == "fail" {
		switch result["message"] {
		case auth.ErrIncorrectUsernameOrEmail.Error():
			c.JSON(http.StatusBadRequest, gin.H{"message": result["message"]})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": result["message"]})
		}
		return
	}

	token := fmt.Sprintf("Bearer %s", result["token"])

	h.infoLogger.Print("successfully logged in")
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": fmt.Sprintf("Welcome %s", user.Username), "Authorization": token})
}
