/* A project by Brian Munger
10/30/2025 */

// SOURCE: https://go.dev/doc/tutorial/web-service-gin

package main

// Import required packages
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

// List of planes, initalize with Cessna 172 :)
var airplanes = []Plane{
	{ID: 1, Model: "C172", Manufacturer: "Cessna", RangeMI: 400., FixedGear: true},
}
var nextID = 2

/* End points */

// Return all planes
func getPlanes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, airplanes)
}

// Return plane by ID
func getPlaneByID(c *gin.Context) {
	id_str := c.Param("id")
	id, err := strconv.Atoi(id_str)

	// Return error (plane ID was in invalid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid plane ID (check format)"})
		return
	}

	// Loop through planes until matching ID is found
	for _, a := range airplanes {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	// Otherwise, return no plane found
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "plane with that ID was not found"})
}

// Add a plane to the dict
func setPlanes(c *gin.Context) {
	var newPlane Plane

	// Return error (error setting the plane)
	if err := c.BindJSON(&newPlane); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the ID, increment for next plane
	newPlane.ID = nextID
	nextID++

	// Add the new plane to the slice
	airplanes = append(airplanes, newPlane)
	c.IndentedJSON(http.StatusCreated, newPlane)
}

// Remove a plane from the dict with the desired ID
func deletePlaneByID(c *gin.Context) {
	id_str := c.Param("id")
	id, err := strconv.Atoi(id_str)

	// Return error (plane ID was in invalid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid plane ID (check format)"})
		return
	}

	// Loop through planes until matching ID is found
	for i, a := range airplanes {
		if a.ID == id {
			/* :i = all the elements before index i
			i+1: = all the elements after index i
			... = expands the slice so append can join */
			airplanes = append(airplanes[:i], airplanes[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "plane deleted!"})
			return
		}
	}

	// Otherwise, return no plane found
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "plane with that ID was not found"})
}

// Main function for routing
func main() {
	router := gin.Default()                           // Initalize a gin router
	router.GET("/planes", getPlanes)                  // curl http://localhost:8000/planes
	router.GET("/planes/:id", getPlaneByID)           // curl http://localhost:8000/planes/1
	router.POST("/planes", setPlanes)                 // curl.exe http://localhost:8000/planes --include --header "Content-Type: application/json" --request "POST" --data '{\"model\":\"B737\",\"manufacturer\":\"Boeing\",\"range_mi\":4000,\"fixed_gear\":false}'
	router.DELETE("/planes/del/:id", deletePlaneByID) // curl.exe -X DELETE http://localhost:8000/planes/del/1

	router.Run("localhost:8000") // Run the local host
}
