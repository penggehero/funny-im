package ws

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	// upgrader  websocket upgrade
	Upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		EnableCompression: true,
	}
)
