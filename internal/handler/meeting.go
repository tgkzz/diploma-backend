package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"server/internal/model"
)

func (h *Handler) makeMeeting(c echo.Context) error {
	userEmail := c.Get("email")

	var req model.MakeAppointmentRequest
	if err := c.Bind(&req); err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	resp, err := h.service.Meeting.CreateMeeting(req, userEmail.(string))
	if err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	response := map[string]interface{}{
		"status":  "success",
		"message": "Successfully created new meeting",
		"roomId":  resp,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) GetMeetingByRoomId(c echo.Context) error {

	roomId := c.QueryParam("room_id")

	meet, err := h.service.Meeting.GetMeetingByRoomId(roomId)
	if err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, meet)
}
