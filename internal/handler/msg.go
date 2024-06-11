package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"server/internal/model"
	"strconv"
)

func (h *Handler) sendMsg(c echo.Context) error {
	var req model.Msg
	if err := c.Bind(&req); err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	if err := h.service.Auth.SendMsg(req.To, req.Msg); err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) GetUserById(c echo.Context) error {
	userId := c.Param("user_id")

	id, err := strconv.Atoi(userId)
	if err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusBadRequest)
	}

	res, err := h.service.Auth.GetUserById(id)
	if err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, res)
}
