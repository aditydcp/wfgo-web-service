package main

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"aditydcp/wfgo-web-service/db"
	"aditydcp/wfgo-web-service/routes"
)
  
func main() {
	var client *mongo.Client
	client = db.ConnectDb()
	routes.Client = client
	router := gin.Default()

	router.GET("/movies", routes.GetMovies)

	defer db.DisconnectDb(client)

	router.Run(":3000")
}