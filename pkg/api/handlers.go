package api

import (
	"log"
	"net/http"
	"time"

	"github.com/IPampurin/WebSocket-Server/pkg/upgrader"
	"github.com/gorilla/websocket"
)

// HandleConnections
func HandleConnections(w http.ResponseWriter, r *http.Request) {

	// получаем экземпляр апгрейдера
	upgrader := upgrader.GetWSUpgrader()

	// меняем тип соединения
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Ошибка при апгрейде: %v", err)
		return
	}
	defer conn.Close() // закрываем соединение при выходе

	log.Printf("Клиент успешно подключен: %s\n", conn.RemoteAddr())

	// поддержка пинг/понга
	conn.SetReadLimit(512) // Максимальный размер сообщения
	conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	// бесконечный цикл для чтения и записи сообщений
	for {
		// читаем сообщение от клиента
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Ошибка чтения: %v\n", err)
			}
			break // выходим из цикла, если клиент отключился или произошла ошибка
		}

		// выводим полученное сообщение в консоль сервера
		log.Printf("Получено сообщение: %s", msg)

		// отправляем сообщение обратно (эхо)
		if err := conn.WriteMessage(messageType, msg); err != nil {
			log.Printf("Ошибка записи: %v", err)
			break
		}
	}

	log.Printf("Клиент отключен: %s\n", conn.RemoteAddr())
}
