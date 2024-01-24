package handler

import (
	"auth/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func (h *Handler) register(c echo.Context) error {
	user := models.User{}

	if err := c.Bind(&user); err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	if err := h.service.Auth.CreateNewUser(user); err != nil {
		h.errorLogger.Print(err)
		if err == models.ErrInvalidEmail || err == models.ErrInvalidPassword {
			return ErrorHandler(c, err, http.StatusBadRequest)
		}
		if strings.Contains(err.Error(), "pq: повторяющееся значение ключа нарушает ограничение уникальности \"users_username_key\"") {
			return ErrorHandler(c, models.ErrUsernameAlreadyTaken, http.StatusBadRequest)
		}
		if strings.Contains(err.Error(), "pq: повторяющееся значение ключа нарушает ограничение уникальности \"users_email_key\"") {
			return ErrorHandler(c, models.ErrEmailAlreadyTaken, http.StatusBadRequest)
		}
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	successResponse := map[string]interface{}{
		"status":  "success",
		"message": "Successfully created new user",
	}
	h.infoLogger.Print("Successfully created new user")
	return c.JSON(http.StatusCreated, successResponse)
}

func (h *Handler) login(c echo.Context) error {
	var creds models.User
	if err := c.Bind(&creds); err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	// TODO: handle errors
	user, err := h.service.Auth.CheckUserCreds(creds)
	if err != nil {
		if err == models.ErrIncorrectUsernameOrEmail || strings.Contains(err.Error(), "sql: no rows in result set") {
			return ErrorHandler(c, models.ErrIncorrectUsernameOrEmail, http.StatusBadRequest)
		}
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	token, err := h.service.Auth.JwtAuthorization(user)
	if err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	response := map[string]interface{}{
		"status":  "success",
		"message": "Successfully logged in",
		"token":   token,
	}

	return c.JSON(http.StatusOK, response)
}
