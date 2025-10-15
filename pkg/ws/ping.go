package ws

import "time"

var (
	pongWait = 10 * time.Second

	pingInterval = (pongWait * 9) / 10
)
