package tests

import (
	"aditydcp/wfgo-web-service/models"
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var BaseURL string
var ServerPort string
var Uri string
var Client *mongo.Client

func TestEnvironment(t *testing.T) {
	assertUri := assert.New(t)
	assertPort := assert.New(t)
	assertUrl := assert.New(t)

	// load environment variables
	godotenv.Load("../.env")

	// connect to mongodb
	Uri := os.Getenv("MONGODB_URI")
	BaseURL := os.Getenv("BASE_URL")
	ServerPort := os.Getenv("SERVER_PORT")

	assertUri.NotEqual(t, Uri, "URI is not set")
	assertPort.NotEqual(t, ServerPort, "", "Port is not set")
	assertUrl.NotEqual(t, BaseURL, "", "Base URL is not set")
}

func TestConnection(t *testing.T) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(Uri))
	assert.NotNil(t, err)
	Client = client
}

func TestAddNewFarm(t *testing.T) {
	assertReq := assert.New(t)
	var emptyV []string
	newFarm := models.Farm{
		ID:      "99",
		Name:    "Test Farm",
		PondIds: emptyV}
	requestUrl := fmt.Sprintf("%s:%s/farm", BaseURL, ServerPort)
	req, errReq := http.NewRequest(http.MethodPost, requestUrl, newFarm, false)
	assertReq.NotNil(errReq)

	assertRes := assert.New(t)
	res, errRes := http.DefaultClient.Do(req)
	assertRes.NotNil(errRes)

	fmt.Printf("%v", res)
}
