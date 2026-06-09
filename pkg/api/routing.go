package api

import "net/http"

func ApiInit() {

	http.HandleFunc("/ws", HandleConnections)
}
