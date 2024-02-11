package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func (h *Handler) getCourseByName(c echo.Context) error {
	name := c.Param("name")

	res, err := h.service.Course.GetCourseByName(name)
	if err != nil {
		h.errorLogger.Print(err)
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			return ErrorHandler(c, err, http.StatusNotFound)
		}
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	successResponse := map[string]interface{}{
		"status":  "success",
		"message": "Successfully got course",
		"course":  res,
	}

	h.infoLogger.Print("Successfully got course")
	return c.JSON(http.StatusOK, successResponse)
}

func (h *Handler) getCourseById(c echo.Context) error {
	id := c.Param("id")

	res, err := h.service.Course.GetCourseById(id)
	if err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	successResponse := map[string]interface{}{
		"status":  "success",
		"message": "Successfully got course",
		"course":  res,
	}

	h.infoLogger.Print("Successfully got course")
	return c.JSON(http.StatusOK, successResponse)
}

func (h *Handler) getAllPost(c echo.Context) error {
	res, err := h.service.Course.GetAllCourse()
	if err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	successResponse := map[string]interface{}{
		"status":  "success",
		"message": "Successfully got all courses",
		"course":  res,
	}

	h.infoLogger.Print("Successfully got course")
	return c.JSON(http.StatusOK, successResponse)
}
