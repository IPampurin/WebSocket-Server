package upgrader

import (
	"github.com/gorilla/websocket"
)

// WSUpgrader инкапсулирует апгрейдер для websocket соединения
type WSUpgrader struct {
	websocket.Upgrader
}

// NewUpgrader возвращает указатель на экземпляр апгрейдера
func NewUpgrader() *WSUpgrader {

	return &WSUpgrader{
		websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}
}
