package server

import (
	"context"
	"log"
	"sync"

	"github.com/redis/go-redis/v9"
)

// ClientManager tracks active WebSocket clients
type ClientManager struct {
	Clients map[string]*Client
	Mutex   sync.Mutex
}

var ctx = context.Background()
var redisClient *redis.Client

// NewClientManager creates a new instance
func NewClientManager() *ClientManager {
	return &ClientManager{
		Clients: make(map[string]*Client),
	}
}

// Initialize Redis Client (called from main)
func InitRedisClient() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,
	})
}

// Save WebSocket connection in Redis
func CacheWsConnection(userID string, addr string) {
	err := redisClient.Set(ctx, "ws:"+userID, addr, 0).Err()
	if err != nil {
		log.Println("Error saving connection:", err)
	}
}

// Get all active WebSocket connections
func GetActiveWsConnections() []string {
	keys, err := redisClient.Keys(ctx, "ws:*").Result()
	if err != nil {
		log.Println("Error retrieving active connections:", err)
		return nil
	}
	return keys
}
