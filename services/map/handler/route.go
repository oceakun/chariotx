package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RouteRequest represents the payload required to compute a route.
// @Description Coordinates for source and destination
type RouteRequest struct {
	SourceLat float64 `json:"source_lat" example:"12.9716"`
	SourceLon float64 `json:"source_lon" example:"77.5946"`
	DestLat   float64 `json:"dest_lat" example:"13.0827"`
	DestLon   float64 `json:"dest_lon" example:"80.2707"`
}

// RegisterRoutes registers the Map Service routes
func RegisterRoutes(router *gin.Engine) {
	router.POST("/route", handleRoute)
}

// handleRoute handles incoming HTTP requests to compute a route.
//
// @Summary Compute Route
// @Description Given source and destination coordinates, returns a route
// @Tags Map
// @Accept json
// @Produce json
// @Param request body RouteRequest true "Source and Destination Coordinates"
// @Success 200 {object} map[string]interface{} "Route computed successfully"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /route [post]
func handleRoute(c *gin.Context) {
	var req RouteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// Forward to GPS Service
	gpsURL := "http://gps-service:8082/compute-route"
	body, _ := json.Marshal(req)
	resp, err := http.Post(gpsURL, "application/json", bytes.NewBuffer(body))
	if err != nil || resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch route from GPS"})
		return
	}
	defer resp.Body.Close()

	responseData, _ := ioutil.ReadAll(resp.Body)
	c.Data(http.StatusOK, "application/json", responseData)
}