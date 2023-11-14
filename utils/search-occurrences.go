package utils

import (
	"aditydcp/wfgo-web-service/db"
	"aditydcp/wfgo-web-service/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

// helper function to Find any instance where Pond Id already taken by other farm
func FindOverPondId(ids []string) ([]models.Farm, []error) {
	var coll = db.Db.Collection("farms")
	cursor, err1 := coll.Find(context.TODO(), bson.M{"ponds": bson.M{"$in": ids}})

	var farms []models.Farm
	var err2 = cursor.All(context.TODO(), &farms)

	var err []error
	err = append(err, err1, err2)

	return farms, err
}

// helper function to Find the existance of a Pond Id
func FindAnyPondId(ids []string) ([]models.Pond, []error) {
	var coll = db.Db.Collection("ponds")
	cursor, err1 := coll.Find(context.TODO(), bson.M{"id": bson.M{"$in": ids}})

	var ponds []models.Pond
	var err2 = cursor.All(context.TODO(), &ponds)

	var err []error
	err = append(err, err1, err2)

	return ponds, err
}

// func FindPondInFarms(id string) ([]models.Farm, []error) {
// 	var coll = db.Db.Collection("farms")
// 	filter := bson.D{{"ponds", bson.D{{"$all", bson.A{id}}}}}
// 	cursor, err1 := coll.Find(context.TODO(), filter)
// 	var err2 error = nil
// 	var results []models.Farm
// 	if err1 != nil {
// 		err2 = cursor.All(context.TODO(), &results)
// 	}

// 	var err []error
// 	err = append(err, err1, err2)

// 	return results, err
// }
