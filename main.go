package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// pond represents data about a pond
type pond struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// farm represents data about a farm
type farm struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// Ponds []pond  `json:"ponds"`
}

// placeholder dummy data
var ponds = []pond{
	{ID: "1", Name: "A"},
	{ID: "2", Name: "B"},
	{ID: "3", Name: "C"},
	{ID: "4", Name: "ABZ"},
}

var farms = []farm{
	{ID: "1", Name: "Blue Train"},
	{ID: "2", Name: "Jeru"},
}

func main() {
	router := gin.Default()
	router.GET("/farms", getFarms)
	router.POST("/farms", addFarm)

	router.Run("localhost:8080")
}

// # REGION START - Handler

// get all farms data
func getFarms(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, farms)
}

// add new farm data
func addFarm(c *gin.Context) {
	var newFarm farm

	// Call BindJSON to bind the received JSON to
	// newFarm.
	if err := c.BindJSON(&newFarm); err != nil {
		return
	}

	// Add to the slice.
	farms = append(farms, newFarm)
	c.IndentedJSON(http.StatusCreated, newFarm)
}
