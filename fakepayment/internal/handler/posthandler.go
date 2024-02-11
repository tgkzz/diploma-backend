package handler

import (
	"fakepayment/internal/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) buyCourse(c echo.Context) error {
	var input model.ClientInput
	if err := c.Bind(&input); err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	if err := h.service.Pay.BuyCourse(input); err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	successResponse := map[string]interface{}{
		"status":  "success",
		"message": "Successfully bought course",
	}

	h.infoLogger.Print("Successfully bought course")
	return c.JSON(http.StatusOK, successResponse)
}
