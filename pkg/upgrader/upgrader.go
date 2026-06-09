package upgrader

import (
	"github.com/gorilla/websocket"
)

type Upgrader struct {
	websocket.Upgrader
}

func NewUpgrader() *Upgrader {

	upgrader := Upgrader{
		websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}

	return &upgrader
}
