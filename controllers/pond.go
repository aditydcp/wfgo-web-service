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

// GET all ponds
func GetPonds(c *gin.Context) {
	AddStats("GET /ponds", c.ClientIP())

	// find all entries
	var coll = db.Db.Collection("ponds")
	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		utils.ErrorReq(c, err)
		return
	}

	// unpack the results
	var ponds []models.Pond
	if err = cursor.All(context.TODO(), &ponds); err != nil {
		utils.ErrorReq(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"count": len(ponds), "result": ponds})
}

// GET pond by id
func GetPondById(c *gin.Context) {
	AddStats("GET /pond/:id", c.ClientIP())

	id := c.Param("id")
	if id == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "No ID given in path"})
		return
	}

	// Look for pond with given ID
	result, err := utils.FindPondId(id)
	if err == mongo.ErrNoDocuments {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "pond data not found"})
		return
	}
	if err != nil {
		utils.ErrorReq(c, err)
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}

// POST create pond
func CreatePond(c *gin.Context) {
	AddStats("POST /pond", c.ClientIP())

	// bind request body into Pond model
	var coll = db.Db.Collection("ponds")
	var newPond models.Pond
	if err := c.BindJSON(&newPond); err != nil {
		utils.ErrorReq(c, err)
		return
	}

	// checks for pond duplicate
	_, err := utils.FindPondId(newPond.ID)
	if err == nil {
		// if id already exists
		c.IndentedJSON(http.StatusConflict,
			gin.H{"message": "pond with id " + newPond.ID + " already exist"})
		return
	}

	_, errQuery := coll.InsertOne(context.TODO(), newPond)
	if errQuery != nil {
		utils.ErrorReq(c, errQuery)
		return
	}

	c.IndentedJSON(http.StatusCreated,
		gin.H{"message": "sucessfully created new pond", "data": newPond})
}

// PUT edit pond
func UpdatePond(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "No ID given in path"})
		return
	}

	// bind request body into Pond model
	var coll = db.Db.Collection("ponds")
	var newPond models.Pond
	if err := c.BindJSON(&newPond); err != nil {
		utils.ErrorReq(c, err)
		return
	}

	// checks for target id pond duplicate
	_, errTargetDupe := utils.FindPondId(newPond.ID)
	if errTargetDupe == nil && id != newPond.ID {
		// if target id already exists
		c.IndentedJSON(http.StatusConflict,
			gin.H{"message": "pond with id " + newPond.ID + " already exist"})
		return
	}

	// checks for pond duplicate from path parameter
	duplicate, errIdDupe := utils.FindPondId(id)
	if errIdDupe == nil {
		// if given id already exists, replace
		_, errQuery := coll.ReplaceOne(context.TODO(),
			bson.D{{"id", id}},
			newPond)
		if errQuery != nil {
			utils.ErrorReq(c, errQuery)
			return
		}
		c.IndentedJSON(http.StatusOK,
			gin.H{"message": "sucessfully update pond",
				"new_data": newPond,
				"old_data": duplicate})
		return
	}
	if errIdDupe != mongo.ErrNoDocuments {
		// if other error
		utils.ErrorReq(c, errIdDupe)
		return
	}

	if id != newPond.ID {
		c.IndentedJSON(http.StatusBadRequest,
			gin.H{"message": "ambiguous ID from path and body. please set the value for ID."})
		return
	}

	// if no entry of given id, insert new
	_, errQuery := coll.InsertOne(context.TODO(), newPond)
	if errQuery != nil {
		utils.ErrorReq(c, errQuery)
		return
	}

	c.IndentedJSON(http.StatusCreated,
		gin.H{"message": "created new pond", "data": newPond})
}
