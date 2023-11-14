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
	AddStats("GET /farms", c.ClientIP())
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
	AddStats("GET /farm/:id", c.ClientIP())
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
	AddStats("POST /pond", c.ClientIP())
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
		// if id already exists
		c.IndentedJSON(http.StatusConflict,
			gin.H{"message": "farm with id " + newFarm.ID + " already exist"})
		return
	}

	if len(newFarm.PondIds) > 0 {
		// checks for pond duplicate
		dupeFarms, errors := utils.FindOverPondId(newFarm.PondIds)
		if errors[0] == nil && len(dupeFarms) > 0 {
			// id already exists
			c.IndentedJSON(http.StatusConflict,
				gin.H{"message": "one or more pond(s) you are going to assign already taken by other farm"})
			return
		}

		// checks for ponds existance
		ponds, errorPonds := utils.FindAnyPondId(newFarm.PondIds)
		if errorPonds[0] == mongo.ErrNoDocuments || len(ponds) == 0 {
			// pond id do not exist
			c.IndentedJSON(http.StatusNotFound,
				gin.H{"message": "one or more pond(s) you are going to assign do(es) not exist"})
			return
		}
	}

	_, errQuery := coll.InsertOne(context.TODO(), newFarm)
	if errQuery != nil {
		utils.ErrorReq(c, errQuery)
		return
	}

	c.IndentedJSON(http.StatusCreated,
		gin.H{"message": "sucessfully created new farm", "data": newFarm})
}

// PUT edit farm
func UpdateFarm(c *gin.Context) {
	AddStats("PUT /farm/:id", c.ClientIP())
	id := c.Param("id")
	if id == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "No ID given in path"})
		return
	}

	// bind request body into Farm model
	var coll = db.Db.Collection("farms")
	var newFarm models.Farm
	if err := c.BindJSON(&newFarm); err != nil {
		utils.ErrorReq(c, err)
		return
	}

	if len(newFarm.PondIds) > 0 {
		// checks for pond duplicate
		_, errors := utils.FindOverPondId(newFarm.PondIds)
		if errors[0] == nil {
			// id already exists
			c.IndentedJSON(http.StatusConflict,
				gin.H{"message": "one or more pond(s) you are going to assign already taken by other farm"})
			return
		}

		// checks for ponds existance
		ponds, errorPonds := utils.FindAnyPondId(newFarm.PondIds)
		if errorPonds[0] == mongo.ErrNoDocuments || len(ponds) == 0 {
			// pond id do not exist
			c.IndentedJSON(http.StatusNotFound,
				gin.H{"message": "one or more pond(s) you are going to assign do(es) not exist"})
			return
		}
	}

	// checks for target id farm duplicate
	_, errTargetDupe := utils.FindFarmId(newFarm.ID)
	if errTargetDupe == nil && id != newFarm.ID {
		// if target id already exists
		c.IndentedJSON(http.StatusConflict,
			gin.H{"message": "farm with id " + newFarm.ID + " already exist"})
		return
	}

	// checks for farm duplicate from path parameter
	duplicate, errIdDupe := utils.FindFarmId(id)
	if errIdDupe == nil {
		// if given id already exists, replace
		_, errQuery := coll.ReplaceOne(context.TODO(),
			bson.D{{"id", id}},
			newFarm)
		if errQuery != nil {
			utils.ErrorReq(c, errQuery)
			return
		}
		c.IndentedJSON(http.StatusOK,
			gin.H{"message": "sucessfully update farm",
				"new_data": newFarm,
				"old_data": duplicate})
		return
	}
	if errIdDupe != mongo.ErrNoDocuments {
		// if other error
		utils.ErrorReq(c, errIdDupe)
		return
	}

	if id != newFarm.ID {
		c.IndentedJSON(http.StatusBadRequest,
			gin.H{"message": "ambiguous ID from path and body. please set the same value for ID."})
		return
	}

	// if no entry of given id, insert new
	_, errQuery := coll.InsertOne(context.TODO(), newFarm)
	if errQuery != nil {
		utils.ErrorReq(c, errQuery)
		return
	}

	c.IndentedJSON(http.StatusCreated,
		gin.H{"message": "created new farm", "data": newFarm})
}
