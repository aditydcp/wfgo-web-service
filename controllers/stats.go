package controllers

import (
	"aditydcp/wfgo-web-service/db"
	"aditydcp/wfgo-web-service/models"
	"aditydcp/wfgo-web-service/utils"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddStats(path string, ip string) {
	var coll = db.Db.Collection("stats")
	// find path
	var result models.Stats
	var err = coll.FindOne(context.TODO(), bson.D{{"path", path}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		// if not found, create entry
		var users = []string{ip}
		var newStat = models.Stats{Path: path, Count: 1, UniqueUser: users}
		coll.InsertOne(context.TODO(), newStat)
		return
	}
	if err != nil {
		return
	}

	// if path stat exists, increment count and append ip if unique
	filter := bson.D{{"path", path}}
	// var update = bson.D{{"$set", bson.D{{"$inc", bson.D{{"count", 1}}}}}}
	var newStat = models.Stats{Path: path, Count: result.Count + 1, UniqueUser: result.UniqueUser}
	var isUnique = true
	for _, x := range result.UniqueUser {
		if ip == x {
			isUnique = false
			break
		}
	}
	if isUnique {
		newUsers := append(result.UniqueUser, ip)
		newStat = models.Stats{Path: path, Count: result.Count + 1, UniqueUser: newUsers}
		// update = bson.D{{"$set", bson.D{{"unique_user_agent", newUsers},
		// 	{"$inc", bson.D{{"count", 1}}}}}}
	}
	coll.UpdateOne(context.TODO(), filter, newStat)
}

// GET all stats
func GetStats(c *gin.Context) {
	// find all entries
	var coll = db.Db.Collection("stats")
	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		utils.ErrorReq(c, err)
		return
	}

	// unpack the results
	var stats []models.Stats
	if err = cursor.All(context.TODO(), &stats); err != nil {
		utils.ErrorReq(c, err)
		return
	}

	var statsPublic []models.StatsPublic
	for i, stat := range stats {
		statsPublic[i] = models.StatsPublic{
			Path: stat.Path, Count: stat.Count, UniqueUserCount: len(stat.UniqueUser)}
	}

	c.JSON(http.StatusOK, gin.H{"result": statsPublic})
}
