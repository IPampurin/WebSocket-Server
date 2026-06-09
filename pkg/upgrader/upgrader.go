package upgrader

import (
	"github.com/gorilla/websocket"
)

// WSUpgrader инкапсулирует апгрейдер для websocket соединения
type WSUpgrader struct {
	websocket.Upgrader
}

// NewWSUpgrader возвращает указатель на экземпляр апгрейдера
func NewWSUpgrader() *WSUpgrader {

	return &WSUpgrader{
		websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}
}
