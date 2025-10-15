package ws

import (
	"encoding/json"
	"log"
	"log/slog"
	"time"

	"github.com/gorilla/websocket"
)

type ClientList = map[string]*Client

type Client struct {
	socketID string
	conn     *websocket.Conn

	manager *Manager
	egress  chan Event
}

func NewClient(socketId string, conn *websocket.Conn, m *Manager) *Client {
	return &Client{
		socketID: socketId,
		conn:     conn,
		manager:  m,
		egress:   make(chan Event),
	}
}

func (c *Client) readMessages() {
	defer func() {
		c.manager.removeClient(c.socketID)
	}()

	if err := c.conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		slog.Error(err.Error())
		return
	}

	c.conn.SetPongHandler(c.pongHandler)

	for {
		_, payload, err := c.conn.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error reading message: %v", err)
			}

			break
		}

		var event Event
		if err := json.Unmarshal(payload, &event); err != nil {
			break
		}

		if err := c.manager.routeEvent(event, c); err != nil {
			slog.Error("error handling message")
		}
	}

}

func (c *Client) broadcast() {
	ticker := time.NewTicker(pingInterval)
	defer func() {
		ticker.Stop()
		c.manager.removeClient(c.socketID)
	}()

	for {
		select {
		case message, ok := <-c.egress:
			if !ok {
				return
			}

			data, err := json.Marshal(message)
			if err != nil {
				return
			}

			if err := c.conn.WriteMessage(websocket.TextMessage, data); err != nil {
				slog.Error(err.Error())
			}
		case <-ticker.C:
			return
		}

	}

}

func (c *Client) pongHandler(pongMsg string) error {
	return c.conn.SetReadDeadline(time.Now().Add(pongWait))
}
