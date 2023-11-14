package routes

import (
	"context"
	"encoding/json"
	"fmt"

	// "aditydcp/wfgo-web-service/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"aditydcp/wfgo-web-service/db"
)

func GetMovies(c *gin.Context) {
	// var (
	// 	farm   models.Farm
	// 	result gin.H
	// )

	coll := db.Client.Database("sample_mflix").Collection("movies")
	title := "Back to the Future"
	var result bson.M
	var err error
	err = coll.FindOne(context.TODO(), bson.D{{"title", title}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", title)
		return
	}
	if err != nil {
		panic(err)
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
}