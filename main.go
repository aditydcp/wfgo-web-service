package main

import (
	"github.com/gin-gonic/gin"

	"aditydcp/wfgo-web-service/models"

	"aditydcp/wfgo-web-service/controllers"
)

// // pond represents data about a pond
// type pond struct {
// 	ID   string `json:"id"`
// 	Name string `json:"name"`
// }

// // farm represents data about a farm
// type farm struct {
// 	ID   string `json:"id"`
// 	Name string `json:"name"`
// 	// Ponds []pond  `json:"ponds"`
// }

// // placeholder dummy data
// var ponds = []pond{
// 	{ID: "1", Name: "A"},
// 	{ID: "2", Name: "B"},
// 	{ID: "3", Name: "C"},
// 	{ID: "4", Name: "ABZ"},
// }

// var farms = []farm{
// 	{ID: "1", Name: "Blue Train"},
// 	{ID: "2", Name: "Jeru"},
// }

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"message": "hello world"})
	// })

	router.GET("/farms", controllers.getFarms)

	// router.GET("/farms", getFarms)
	// router.GET("/farms/:id", getFarmByID)
	// router.POST("/farms", addFarm)

	router.Run("localhost:8080")
}

// // # REGION START - Handler

// // get all farms data
// func getFarms(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, farms)
// }

// // add new farm data
// func addFarm(c *gin.Context) {
// 	var newFarm farm

// 	// Call BindJSON to bind the received JSON to
// 	// newFarm.
// 	if err := c.BindJSON(&newFarm); err != nil {
// 		return
// 	}

// 	// Add to the slice.
// 	farms = append(farms, newFarm)
// 	c.IndentedJSON(http.StatusCreated, newFarm)
// }

// // get farm data by id
// func getFarmByID(c *gin.Context) {
// 	id := c.Param("id")

// 	// Look for farm with given ID
// 	for _, a := range farms {
// 		if a.ID == id {
// 			c.IndentedJSON(http.StatusOK, a)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "farm not found"})
// }
