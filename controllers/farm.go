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

// GET all farms
func GetFarms(c *gin.Context) {
	// find all entries
	var coll = db.Db.Collection("farms")
	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		utils.ErrorReq(c, err)
		return
	}

	// unpack the results
	var farms []models.Farm
	if err = cursor.All(context.TODO(), &farms); err != nil {
		utils.ErrorReq(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"count": len(farms), "result": farms})
}

// GET farm by id
func GetFarmById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "No ID given in path"})
		return
	}

	// Look for farm with given ID
	result, err := utils.FindFarmId(id)
	if err == mongo.ErrNoDocuments {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "farm data not found"})
		return
	}
	if err != nil {
		utils.ErrorReq(c, err)
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}

// POST create farm
func CreateFarm(c *gin.Context) {
	// bind request body into Farm model
	var coll = db.Db.Collection("farms")
	var newFarm models.Farm
	if err := c.BindJSON(&newFarm); err != nil {
		utils.ErrorReq(c, err)
		return
	}

	// checks for farm duplicate
	_, err := utils.FindFarmId(newFarm.ID)
	if err == nil {
		c.IndentedJSON(http.StatusConflict,
			gin.H{"message": "farm with id " + newFarm.ID + " already exist"})
		return
	}

	_, errQuery := coll.InsertOne(context.TODO(), newFarm)
	if errQuery != nil {
		utils.ErrorReq(c, errQuery)
		return
	}

	c.IndentedJSON(http.StatusCreated,
		gin.H{"message": "sucessfully created new farm", "data": newFarm})
}
