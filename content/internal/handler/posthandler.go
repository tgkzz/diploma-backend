package handler

import (
	"content/internal/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) createCourse(c echo.Context) error {
	var course model.Course

	if err := c.Bind(&course); err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	if err := h.service.Course.CreateNewCourse(course); err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	successResponse := map[string]interface{}{
		"status":  "success",
		"message": "Successfully created new course",
	}

	h.infoLogger.Print("Successfully created new course")
	return c.JSON(http.StatusCreated, successResponse)
}
