package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Plane struct
type Plane struct {
	ID           int     `json:"id"`
	Model        string  `json:"model"`
	Manufacturer string  `json:"manufacturer"`
	RangeMI      float64 `json:"range_mi"`
	FixedGear    bool    `json:"fixed_gear"`
}

// List of planes
var airplanes = []Plane{
	{ID: 1, Model: "C172", Manufacturer: "Cessna", RangeMI: 400., FixedGear: true},
}
var nextID = 2

/* End points */
// Return all planes
func getPlanes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, airplanes)
}

func getPlaneByID(c *gin.Context) {
	id_str := c.Param("id")
	id, err := strconv.Atoi(id_str)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid plane ID"})
		return
	}

	for _, a := range airplanes {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "plane  not found"})
}
func setPlanes(c *gin.Context) {
	var newPlane Plane

	if err := c.BindJSON(&newPlane); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newPlane.ID = nextID
	nextID++

	// Add the new plane to the slice
	airplanes = append(airplanes, newPlane)
	c.IndentedJSON(http.StatusCreated, newPlane)
}

// Main function for routing
func main() {
	router := gin.Default()                 // Initalize a gin router
	router.GET("/planes", getPlanes)        // curl http://localhost:8000/planes
	router.GET("/planes/:id", getPlaneByID) // curl http://localhost:8000/planes/1
	router.POST("/planes", setPlanes)       // curl.exe http://localhost:8000/planes --include --header "Content-Type: application/json" --request "POST" --data '{\"model\":\"B737\",\"manufacturer\":\"Boeing\",\"range_mi\":4000,\"fixed_gear\":false}'

	router.Run("localhost:8000") // Run the local host
}
