package handler

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"net/http"
	"server/internal/model"
	"time"
)

const pingPeriod = 30 * time.Second

type Room struct {
	RoomID    string
	Clients   map[*websocket.Conn]bool
	Broadcast chan Message
	ErrChan   chan error
}

var Rooms = make(map[string]*Room)

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
	roomID := c.Param("room_id")
	conn, err := Upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		h.errorLogger.Print(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer func(conn *websocket.Conn) {
		if err := conn.Close(); err != nil {
			return
		}
	}(conn)

	meeting, err := h.service.Meeting.GetMeetingByRoomId(roomID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrNoMeeting)
	}

	if Rooms[roomID] == nil {
		Rooms[roomID] = &Room{
			RoomID:    meeting.RoomId,
			Clients:   make(map[*websocket.Conn]bool),
			Broadcast: make(chan Message),
			ErrChan:   make(chan error),
		}
		go h.HandleMessages(meeting.RoomId, Rooms[roomID].ErrChan)
	}

	room := Rooms[roomID]
	room.Clients[conn] = true

	for {
		var msg Message
		if err := conn.ReadJSON(&msg); err != nil {
			h.errorLogger.Printf("Error reading JSON: %v", err)
			delete(room.Clients, conn)
			if len(room.Clients) == 0 {
				close(room.Broadcast)
				delete(Rooms, roomID)
			}
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		h.infoLogger.Printf("Received message: %v", msg)
		room.Broadcast <- msg
	}
}

func (h *Handler) HandleMessages(roomID string, errChan chan error) {
	room := Rooms[roomID]
	pingTicker := time.NewTicker(pingPeriod)
	defer pingTicker.Stop()

	for {
		select {
		case msg := <-room.Broadcast:
			for client := range room.Clients {
				if err := client.WriteJSON(msg); err != nil {
					client.Close()
					delete(room.Clients, client)
					if len(room.Clients) == 0 {
						close(room.Broadcast)
						delete(Rooms, roomID)
					}
					errChan <- err
					return
				}
			}
		case <-pingTicker.C:
			for client := range room.Clients {
				if err := client.WriteMessage(websocket.PingMessage, nil); err != nil {
					client.Close()
					delete(room.Clients, client)
				}
			}
		}
	}
}
