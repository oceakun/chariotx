package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
	"os"
	"github.com/coder/websocket"
)

// Client represents a WebSocket connection
type Client struct {
	ID   string
	Conn *websocket.Conn
	IP   string
	Send chan []byte
}

// WebSocketServer manages WebSocket clients and messages
type WebSocketServer struct {
	Clients    map[string]*Client
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
	Mutex      sync.Mutex
}

// NewWebSocketServer initializes the WebSocket server
func NewWebSocketServer() *WebSocketServer {
	return &WebSocketServer{
		Clients:    make(map[string]*Client),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

// HandleConnections upgrades HTTP to WebSocket and manages clients
func (ws *WebSocketServer) HandleConnections(w http.ResponseWriter, r *http.Request) {
	// Accept WebSocket connection
	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		CompressionMode: websocket.CompressionDisabled,
	})
	if err != nil {
		log.Println("Failed to accept WebSocket connection:", err)
		return
	}

	clientIP := r.RemoteAddr // Extract client IP
	log.Println("Client connected:", clientIP)

	clientID := fmt.Sprintf("%p", conn)
	client := &Client{
		ID:   clientID,
		Conn: conn,
		IP:   clientIP,
		Send: make(chan []byte),
	}

	// Save connection info in Redis
	CacheWsConnection(clientID, clientIP)

	// Register new client
	ws.Register <- client

	// Handle messages
	go ws.readMessages(client)
	go ws.writeMessages(client)
}


// Read and parse JSON message
func ReadLocation(conn *websocket.Conn) (*LocationData, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, msg, err := conn.Read(ctx)
	if err != nil {
		return nil, err
	}

	var location LocationData
	if err := json.Unmarshal(msg, &location); err != nil {
		return nil, err
	}

	return &location, nil
}

// readMessages listens for incoming messages
func (ws *WebSocketServer) readMessages(client *Client) {
	defer func() {
		ws.Unregister <- client
		client.Conn.Close(websocket.StatusNormalClosure, "Closing connection")
	}()

	for {
		location, err := ReadLocation(client.Conn)
		if err != nil {
			log.Printf("Error reading message from client %s: %v", client.ID, err)
			break
		}

		// Store location locally
		StoreLocationToFile(location)
		// instead of storing locally, send the data to a Cassandra instance through the Location service

		// Broadcast to all clients
		// ws.Broadcast <- []byte(fmt.Sprintf("User %s was at (%f, %f) at (%f)", location.UserID, location.Lat, location.Lng, location.Timestamp))
	}
}

// StoreLocationToFile logs locations in a file
func StoreLocationToFile(location *LocationData) {
	filePath := "user_locations.log"
	entry := fmt.Sprintf("%s: (%f, %f, %f)\n", location.UserID, location.Lat, location.Lng, location.Timestamp)

	// Use os.OpenFile instead of undefined openFile
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	_, err = f.WriteString(entry)
	if err != nil {
		log.Println("Error writing to file:", err)
	}
}

// writeMessages sends messages to the client
func (ws *WebSocketServer) writeMessages(client *Client) {
	for msg := range client.Send {
		err := client.Conn.Write(context.Background(), websocket.MessageText, msg)
		if err != nil {
			log.Printf("Error sending message to client %s: %v", client.ID, err)
			break
		}
	}
}

// Run manages client connections and message broadcasting
func (ws *WebSocketServer) Run() {
	for {
		select {
		case client := <-ws.Register:
			ws.Mutex.Lock()
			ws.Clients[client.ID] = client
			ws.Mutex.Unlock()
			log.Printf("Client %s connected", client.ID)

		case client := <-ws.Unregister:
			ws.Mutex.Lock()
			delete(ws.Clients, client.ID)
			close(client.Send)
			ws.Mutex.Unlock()
			log.Printf("Client %s disconnected", client.ID)

		case msg := <-ws.Broadcast:
			ws.Mutex.Lock()
			for _, client := range ws.Clients {
				client.Send <- msg
			}
			ws.Mutex.Unlock()
		}
	}
}
