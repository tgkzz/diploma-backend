package handler

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"net/http"
)

var Clients = make(map[*websocket.Conn]bool)
var Broadcast = make(chan Message)

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) handleConnections(c echo.Context) error {
	conn, err := Upgrader.Upgrade(c.Response(), c.Request(), c.Response().Header())
	if err != nil {
		h.errorLogger.Print(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer conn.Close()

	Clients[conn] = true

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			h.errorLogger.Print(err)
			delete(Clients, conn)
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		Broadcast <- msg
	}
}

func (h *Handler) HandleMessages() error {

	for {
		msg := <-Broadcast

		for client := range Clients {
			err := client.WriteJSON(msg)
			if err != nil {
				h.errorLogger.Print(err)
				client.Close()
				delete(Clients, client)
			}
		}
	}
}
