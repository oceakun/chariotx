// this test mocks a real user connecting to the service through websockets, sending continuous pings consisting of coordinates (longitude,latitude), their ip address and timestamp for that moment

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/coder/websocket"
	"github.com/google/uuid"
)

// LocationData represents mock location data
type LocationData struct {
	UserID        string  `json:"user_id"`
	Lat       float64 `json:"lat"`
	Lng       float64 `json:"lng"`
	Timestamp float64 `json:"tmstmp"`
}

func main() {
	serverURL := "ws://localhost:8082/ws"

	// Set up signal catching for graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Connect to WebSocket server
	conn, _, err := websocket.Dial(context.Background(), serverURL, nil)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close(websocket.StatusNormalClosure, "Closing mock client")
	log.Println("Connected to WebSocket server")

	// Mock location data with a dynamic timestamp
	mockLocation := LocationData{
		UserID:  uuid.New().String(),
		Lat: 40.7128,
		Lng: -74.0060,
		Timestamp: float64(time.Now().Unix()), // Get the current timestamp
	}

	// Send data every 5 seconds (to keep connection alive)
	go func() {
		for {
			select {
			case <-stop:
				return // Gracefully stop the go routine on receiving signal
			default:
				// Update timestamp each time we send data
				mockLocation.Timestamp = float64(time.Now().Unix())
				data, err := json.Marshal(mockLocation)
				if err != nil {
					log.Println("JSON marshal error:", err)
					return
				}

				err = conn.Write(context.Background(), websocket.MessageText, data)
				if err != nil {
					log.Println("Send error:", err)
					return
				}
				fmt.Println("Sent:", string(data))
				time.Sleep(5 * time.Second) // Keep sending data every 5 seconds
			}
		}
	}()

	// Keep listening for responses from server
	go func() {
		for {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			_, msg, err := conn.Read(ctx)
			cancel()
			if err != nil {
				log.Println("Read error:", err)
				break
			}
			fmt.Println("Received:", string(msg))
		}
	}()

	// Wait for interrupt signal (Ctrl+C)
	<-stop
	log.Println("Mock client exiting...")
}
