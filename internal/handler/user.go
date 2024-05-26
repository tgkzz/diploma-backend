package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"server/internal/model"
	"strings"
)

func (h *Handler) register(c echo.Context) error {
	user := model.User{}

	if err := c.Bind(&user); err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	if err := h.service.Auth.CreateNewUser(user); err != nil {
		h.errorLogger.Print(err)
		if err == model.ErrInvalidEmail || err == model.ErrInvalidPassword || err == model.ErrInvalidName {
			return ErrorHandler(c, err, http.StatusBadRequest)
		}
		if strings.Contains(err.Error(), "pq: повторяющееся значение ключа нарушает ограничение уникальности \"users_email_key\"") {
			return ErrorHandler(c, model.ErrEmailAlreadyTaken, http.StatusBadRequest)
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
	var creds model.User
	if err := c.Bind(&creds); err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	// TODO: handle errors
	user, err := h.service.Auth.CheckUserCreds(creds)
	if err != nil {
		h.errorLogger.Print(err)
		if err == model.ErrIncorrectEmailOrPassword || strings.Contains(err.Error(), "sql: no rows in result set") {
			return ErrorHandler(c, model.ErrIncorrectEmailOrPassword, http.StatusBadRequest)
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
		"fname":   user.FirstName,
		"lname":   user.LastName,
		"email":   user.Email,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) getUserByEmail(c echo.Context) error {
	email := c.Get("email")

	res, err := h.service.Auth.GetUserByEmail(email.(string))
	if err != nil {
		h.errorLogger.Print(err)
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			return ErrorHandler(c, err, http.StatusNotFound)
		}
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	response := map[string]interface{}{
		"status":  "success",
		"message": "Successfully got user by email",
		"user":    res,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) sendEmailCode(c echo.Context) error {
	var req model.SendEmailCodeRequest
	if err := c.Bind(&req); err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	if err := h.service.Auth.SendEmailCode(req.Email, c.Request().Context()); err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	response := map[string]interface{}{
		"status":  "success",
		"message": "Successfully send code",
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) checkEmailCode(c echo.Context) error {
	var req model.CheckEmailCodeRequest
	if err := c.Bind(&req); err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	if err := h.service.Auth.CheckCode(req.Email, req.Code, c.Request().Context()); err != nil {
		h.errorLogger.Print(err)
		if errors.Is(err, model.ErrIncorrectCode) {
			return ErrorHandler(c, err, http.StatusBadRequest)
		}
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	return c.JSON(http.StatusNoContent, nil)
}

func (h *Handler) deleteUser(c echo.Context) error {
	email := c.Get("email")

	if err := h.service.Auth.DeleteUserByEmail(email.(string)); err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) UpdateUser(c echo.Context) error {
	email := c.Get("email")

	var req model.UpdateUserRequest
	if err := c.Bind(&req); err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	if err := h.service.Auth.UpdateUserByEmail(email.(string), req); err != nil {
		h.errorLogger.Print(err)
		if errors.Is(err, model.ErrEmailIsAlreadyUser) {
			return ErrorHandler(c, err, http.StatusBadRequest)
		}
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) buyCourse(c echo.Context) error {
	email := c.Get("email")

	courseId := c.Param("course_id")

	if err := h.service.Course.BuyCourse(c.Request().Context(), courseId, email.(string)); err != nil {
		h.errorLogger.Print(err)
		if errors.Is(err, model.ErrNotFound) || errors.Is(err, model.ErrCourseAlreadyPurchased) {
			return ErrorHandler(c, err, http.StatusBadRequest)
		}
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) getUserCourse(c echo.Context) error {
	courseId := c.Param("course_id")

	result, err := h.service.Course.GetCourse(c.Request().Context(), courseId)
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

func (h *Handler) getUserCourses(c echo.Context) error {
	email := c.Get("email")

	res, err := h.service.Course.GetUserCourses(c.Request().Context(), email.(string))
	if err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	response := map[string]interface{}{
		"status": "success",
		"course": res,
	}

	return c.JSON(http.StatusOK, response)
}
