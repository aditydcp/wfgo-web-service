package main

import (
	"github.com/gin-gonic/gin"
	"aditydcp/wfgo-web-service/db"
	"aditydcp/wfgo-web-service/routes"
	"aditydcp/wfgo-web-service/controllers"
)
  
func main() {
	db.ConnectDb()
	router := gin.Default()

	router.GET("/movies", routes.GetMovies)

	router.GET("/farms", controllers.GetFarms)
	router.GET("/farm/:id", controllers.GetFarmById)
	router.POST("/farm", controllers.CreateFarm)

	defer db.DisconnectDb()

	router.Run(":3000")
}