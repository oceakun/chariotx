package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"github.com/oceakun/chariotx/services/graph-processing/models"
	"github.com/oceakun/chariotx/services/graph-processing/graph"
	"net/http"
)

type RouteRequest struct {
	Source      string `json:"source" binding:"required"`
	Destination string `json:"destination" binding:"required"`
}

func RegisterRoutes(r *gin.Engine, session *gocql.Session) {
	r.POST("/route", func(c *gin.Context) {
		var req RouteRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		segments := []models.Segment{}
		iter := session.Query("SELECT id, source, target, distance FROM road_segments").Iter()

		var seg models.Segment
		for iter.Scan(&seg.ID, &seg.Source, &seg.Target, &seg.Distance) {
			segments = append(segments, seg)
		}
		if err := iter.Close(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
			return
		}

		g := graph.BuildGraph(segments)
		path, distance := graph.Dijkstra(g, req.Source, req.Destination)

		if path == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "No route found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"path":     path,
			"distance": distance,
		})
	})
}