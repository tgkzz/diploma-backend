package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"server/internal/model"
)

func (h *Handler) GetAllCourses(c echo.Context) error {
	result, err := h.service.Course.GetAllCourses(c.Request().Context())
	if err != nil {
		h.errorLogger.Print(err)
		if errors.Is(err, model.ErrNotFound) {
			return ErrorHandler(c, err, http.StatusBadRequest)
		}
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	response := map[string]interface{}{
		"status":  "success",
		"courses": result,
	}

	return c.JSON(http.StatusOK, response)
}
