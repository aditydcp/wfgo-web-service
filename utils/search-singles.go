package utils

import (
	"aditydcp/wfgo-web-service/db"
	"aditydcp/wfgo-web-service/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

// helper function FindOne for farm
func FindFarmId(id string) (models.Farm, error) {
	var coll = db.Db.Collection("farms")
	var result models.Farm
	var err = coll.FindOne(context.TODO(), bson.D{{"id", id}}).Decode(&result)

	return result, err
}

// helper function FindOne for pond
func FindPondId(id string) (models.Pond, error) {
	var coll = db.Db.Collection("ponds")
	var result models.Pond
	var err = coll.FindOne(context.TODO(), bson.D{{"id", id}}).Decode(&result)

	return result, err
}
