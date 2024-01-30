package handler

import (
	"auth/internal/models"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) registerAdmin(c echo.Context) error {
	var admin models.Admin
	if err := c.Bind(&admin); err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	if err := h.service.AdminAuth.CreateNewAdmin(admin); err != nil {
		h.errorLogger.Print(err)
		if errors.Is(err, models.ErrEmptyness) || errors.Is(err, models.ErrInvalidPassword) {
			return ErrorHandler(c, err, http.StatusBadRequest)
		}
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	successResponse := map[string]interface{}{
		"status":  "success",
		"message": "Successfully created new admin",
	}
	h.infoLogger.Print("Successfully created new admin")
	return c.JSON(http.StatusCreated, successResponse)
}

func (h *Handler) loginAdmin(c echo.Context) error {
	var creds models.Admin
	if err := c.Bind(&creds); err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	admin, err := h.service.AdminAuth.CheckAdminCreds(creds)
	if err != nil {
		h.errorLogger.Print(err)
		if errors.Is(err, models.ErrIncorrectUsernameOrPassword) {
			return ErrorHandler(c, err, http.StatusNotFound)
		}
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	token, err := h.service.AdminAuth.JwtAdminAuthorization(admin)
	if err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	response := map[string]interface{}{
		"status":  "success",
		"message": "Successfully logged as admin",
		"token":   token,
	}

	return c.JSON(http.StatusOK, response)
}
