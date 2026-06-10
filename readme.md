# WebSocket Server - от HTTP к WebSocket

### 📋 Описание проекта  

В данном репозитории приведён пример WebSocket-сервера на Go с эхо-обработкой сообщений.  
HTTP-сервер принимает входящие запросы и «поднимает» их до уровня WebSocket.  

### 🖥️ Возможности  

- Преобразование HTTP → WebSocket с помощью `gorilla/websocket`  
- Эхо-ответ на каждое сообщение клиента  
- Graceful shutdown при получении SIGINT/SIGTERM  
- Пинг/понг обработка для поддержания соединения  

### 🗂️ Структура проекта  

```bash
.

├── pkg/
│   ├── api/               
│   │   ├── handlers.go    # логика WebSocket соединения
│   │   └── routing.go     # маршрутизация
│   ├── server/            
│   │   └── server.go      # запуск и graceful shutdown
│   └── upgrader/          
│       └── upgrader.go    # WebSocket апгрейдер (синглтон)
├── main.go                # точка входа, обработка сигналов
├── go.mod
├── go.sum
└── readme.md              # этот файл
```


### 🚀 Быстрый старт  

**Требования:**  

- Go 1.21+  
- Свободный порт: 8081    

**Запуск:**  

    cd ./ && go run main.go  

- Сервер запустится на localhost:8081 (WebSocket endpoint: ws://localhost:8081/ws)  

### 🧪 Тестирование

Через Postman:

    Создайте WebSocket Request  
    URL: ws://localhost:8081/ws  
    Нажмите Connect  
    Отправьте сообщение, получите эхо-ответ  