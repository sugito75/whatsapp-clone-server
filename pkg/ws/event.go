package ws

import "encoding/json"

type EventType string

type EventHandler func(e Event, c *Client) error

type Event struct {
	Type EventType `json:"type"`

	Payload json.RawMessage `json:"payload"`
}

const (
	PrivateMessage EventType = "private_message"
	GroupMessage   EventType = "group_message"
	NewMessage     EventType = "new_message"
	Typing         EventType = "typing"
	NewStatus      EventType = "new_status"
)
