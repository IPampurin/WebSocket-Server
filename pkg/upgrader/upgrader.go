package upgrader

import (
	"github.com/gorilla/websocket"
)

// WSUpgrader инкапсулирует апгрейдер для websocket соединения
type WSUpgrader struct {
	websocket.Upgrader
}

// wsUpgrader единственный экземпляр апгрейдера
var wsUpgrader = WSUpgrader{
	websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	},
}

// GetWSUpgrader возвращает указатель на апгрейдер
func GetWSUpgrader() *WSUpgrader {

	return &wsUpgrader
}
