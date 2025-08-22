package main

import (
	"github.com/gin-gonic/gin"
	"github.com/oceakun/chariotx/services/graph-processing/config"
	"github.com/oceakun/chariotx/services/graph-processing/cassandra"
	"github.com/oceakun/chariotx/services/graph-processing/handler"
)

func main() {
	// Initialize config
	cfg := config.Load()

	// Connect to Cassandra
	session := cassandra.Connect(cfg.Cassandra)
	defer session.Close()

	// Initialize Gin router
	router := gin.Default()

	// Register route handler
	handler.RegisterRoutes(router, session)

	// Start server
	router.Run(":" + cfg.Port)
}