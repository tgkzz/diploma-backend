package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"server/internal/model"
)

func (h *Handler) GetAllCourses(c echo.Context) error {
	result, err := h.service.Course.GetCoursesLimited(c.Request().Context())
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

func (h *Handler) GetCourse(c echo.Context) error {
	courseId := c.Param("course_id")

	result, err := h.service.Course.GetCourseLimited(c.Request().Context(), courseId)
	if err != nil {
		h.errorLogger.Print(err)
		if errors.Is(err, model.ErrNotFound) {
			return ErrorHandler(c, err, http.StatusBadRequest)
		}
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	response := map[string]interface{}{
		"status": "success",
		"course": result,
	}

	return c.JSON(http.StatusOK, response)
}
