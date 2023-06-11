package websocket

import (
	"net/http"
	"vm-backend/configs"

	"github.com/gorilla/websocket"
)

type Server struct {
	*websocket.Upgrader
	configs.WebSocketConfig
}

func New(cfg configs.WebSocketConfig) *Server {
	var upgrader = &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true // Allow all origins
		},
	}

	return &Server{
		upgrader,
		cfg,
	}
}
