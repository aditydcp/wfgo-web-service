package db

import (
	"context"
	"log"
	"os"
	
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectDb() {
	// load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// connect to mongodb
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	Client = client
}

func DisconnectDb() {
	if err := Client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}