package main

import (
	"aditydcp/wfgo-web-service/controllers"
	"aditydcp/wfgo-web-service/db"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		log.Fatal("You must set your 'SERVER_PORT' environment variable.")
	}
	db.ConnectDb()
	router := gin.Default()

	// router.GET("/movies", routes.GetMovies)

	router.POST("/farm", controllers.CreateFarm)
	router.GET("/farms", controllers.GetFarms)
	router.GET("/farm/:id", controllers.GetFarmById)
	router.PUT("/farm/:id", controllers.UpdateFarm)
	router.DELETE("/farm/:id", controllers.DeleteFarm)

	router.POST("/pond", controllers.CreatePond)
	router.GET("/ponds", controllers.GetPonds)
	router.GET("/pond/:id", controllers.GetPondById)
	router.PUT("/pond/:id", controllers.UpdatePond)
	router.DELETE("/pond/:id", controllers.DeletePond)

	router.GET("/recycled/farms", controllers.GetRecycledFarms)
	router.GET("/recycled/ponds", controllers.GetRecycledPonds)

	router.GET("/statistics", controllers.GetStats)

	defer db.DisconnectDb()

	router.Run(fmt.Sprintf(":%s", serverPort))
}
