package message

type MessageDTO struct {
	To      string  `json:"to"`
	Message Message `json:"message"`
}
