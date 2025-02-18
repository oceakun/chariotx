package main

import (
	"log"
	"net/http"
	"websocket-manager/internal/server"
	"websocket-manager/pkg/config"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize Redis
	server.InitRedisClient()

	// Initialize WebSocket server
	wsServer := server.NewWebSocketServer()
	go wsServer.Run()

	// HTTP handler for WebSocket
	http.HandleFunc("/ws", wsServer.HandleConnections)

	log.Printf("WebSocket Manager running on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}
