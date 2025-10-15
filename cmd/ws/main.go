package main

import (
	"log/slog"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/sugito75/chat-app-server/pkg/logger"
	"github.com/sugito75/chat-app-server/pkg/ws"
)

func main() {
	godotenv.Load()
	logger.InitLogger()

	manager := ws.NewManager()

	http.HandleFunc("/ws", manager.HandleConn)

	slog.Info("Socket server listening at", "port", ":4000")
	http.ListenAndServe(":4000", nil)
}
