package ws

import (
	"errors"
	"net/http"
	"sync"

	"github.com/gofiber/fiber/v2/log"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin:     func(r *http.Request) bool { return true },
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type Manager struct {
	clients ClientList
	on      map[EventType]EventHandler
	handler MessageHandler
	sync.RWMutex
}

func NewManager() *Manager {
	m := &Manager{
		clients: make(ClientList),
		on:      make(map[EventType]EventHandler),
		// handler: ,
	}

	m.setupHandlers()
	return m
}

func (m *Manager) HandleConn(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("socketId")

	if id == "" {
		log.Error("no socket id provided")
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error(err)
		return
	}

	client := NewClient(id, conn, m)
	m.addClient(client)

	go client.readMessages()
	go client.broadcast()

}

func (m *Manager) setupHandlers() {
	// m.on[PrivateMessage] = m.handler.HandlePrivateMessage
	// m.on[GroupMessage] = m.handler.HandleGroupMessage
}

func (m *Manager) addClient(c *Client) {
	m.Lock()
	defer m.Unlock()

	m.clients[c.socketID] = c
}

func (m *Manager) removeClient(id string) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.clients[id]; ok {
		m.clients[id].conn.Close()
		delete(m.clients, id)
	}
}

func (m *Manager) routeEvent(e Event, c *Client) error {

	if handler, exists := m.on[e.Type]; exists {
		if err := handler(e, c); err != nil {
			return err
		}
		return nil
	}

	return errors.New("Event not supported")

}
