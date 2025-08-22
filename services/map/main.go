package main

import (
	"github.com/gin-gonic/gin"
	"github.com/oceakun/chariotx/services/map/handler"
)

func main() {
	router := gin.Default()
	handler.RegisterRoutes(router)
	router.Run(":8081")
}