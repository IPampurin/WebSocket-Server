package api

import "net/http"

func RoutingInit() {

	http.HandleFunc("/ws", HandleConnections)
}
