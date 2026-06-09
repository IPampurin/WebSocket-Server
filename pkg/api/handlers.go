package api

import (
	"log"
	"net/http"

	"github.com/IPampurin/WebSocket-Server/pkg/upgrader"
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

	log.Println("Клиент успешно подключен")

	// бесконечный цикл для чтения и записи сообщений
	for {
		// читаем сообщение от клиента
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Ошибка при чтении: %v", err)
			break // выходим из цикла, если клиент отключился или произошла ошибка
		}

		// выводим полученное сообщение в консоль сервера
		log.Printf("Получено сообщение: %s", p)

		// отправляем сообщение обратно (эхо)
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Printf("Ошибка при записи: %v", err)
			break
		}
	}
}
