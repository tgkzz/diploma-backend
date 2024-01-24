package handler

import "github.com/labstack/echo/v4"

func ErrorHandler(c echo.Context, err error, code int) error {
	errResponse := map[string]string{
		"status":  "fail",
		"message": err.Error(),
	}

	return c.JSON(code, errResponse)
}
