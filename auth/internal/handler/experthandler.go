package handler

import (
	"auth/internal/models"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) registerExpert(c echo.Context) error {
	var expert models.Expert
	if err := c.Bind(&expert); err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	if err := h.service.ExpertAuth.CreateExpert(expert); err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	successResponse := map[string]interface{}{
		"status":  "success",
		"message": "Successfully created new expert",
	}
	h.infoLogger.Print("Successfully created new expert")
	return c.JSON(http.StatusCreated, successResponse)
}

func (h *Handler) loginExpert(c echo.Context) error {
	var creds models.Expert
	if err := c.Bind(&creds); err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	exp, err := h.service.ExpertAuth.CheckExpertCreds(creds)
	if err != nil {
		h.errorLogger.Print(err)
		if errors.Is(err, models.ErrIncorrectUsernameOrPassword) {
			return ErrorHandler(c, err, http.StatusNotFound)
		}
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	token, err := h.service.ExpertAuth.JwtExpertAuthorization(exp)
	if err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	response := map[string]interface{}{
		"status":  "success",
		"message": "Successfully logged as expert",
		"token":   token,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) getAllExperts(c echo.Context) error {
	res, err := h.service.ExpertAuth.GetAllExperts()
	if err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	response := map[string]interface{}{
		"status":  "success",
		"message": "Successfully got all experts",
		"experts": res,
	}

	return c.JSON(http.StatusOK, response)
}
