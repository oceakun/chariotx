package server

import (
	"context"
	"encoding/json"
	"log"

	"github.com/coder/websocket"
)

// LocationData struct to store latitude and longitude
type LocationData struct {
	IP        string  `json:"ip"`
	Lat       float64 `json:"lat"`
	Lng       float64 `json:"lng"`
	Timestamp float64 `json:"tmstmp"`
}


// Broadcast location updates to all clients
func BroadcastLocationUpdate(msg LocationData, clients map[string]*websocket.Conn) {
	data, _ := json.Marshal(msg)

	for userID, conn := range clients {
		err := conn.Write(context.Background(), websocket.MessageText, data)
		if err != nil {
			log.Println("Error sending message to", userID, ":", err)
			conn.Close(websocket.StatusInternalError, "Connection closed")
			delete(clients, userID)
		}
	}
}
