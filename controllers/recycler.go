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

// DELETE farm by id
func DeleteFarm(c *gin.Context) {
	AddStats("DELETE /farm/:id", c.ClientIP())
	var coll = db.Db.Collection("farms")
	id := c.Param("id")
	if id == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "No ID given in path"})
		return
	}

	// Look for farm with given ID
	farm, errFind := utils.FindFarmId(id)
	if errFind == mongo.ErrNoDocuments {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "farm data not found"})
		return
	}
	if errFind != nil {
		utils.ErrorReq(c, errFind)
		return
	}

	var recycleColl = db.Db.Collection("deleted_farms")
	_, errQuery := recycleColl.InsertOne(context.TODO(), farm)
	if errQuery != nil {
		utils.ErrorReq(c, errQuery)
		return
	}
	_, errDel := coll.DeleteOne(context.TODO(), bson.D{{"id", id}})
	if errDel != nil {
		utils.ErrorReq(c, errQuery)
		return
	}

	c.IndentedJSON(http.StatusOK,
		gin.H{"message": "data succesfully deleted"})
}

// DELETE pond by id
func DeletePond(c *gin.Context) {
	AddStats("DELETE /pond/:id", c.ClientIP())
	var coll = db.Db.Collection("ponds")
	id := c.Param("id")
	if id == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "No ID given in path"})
		return
	}

	// Look for pond with given ID
	pond, errFind := utils.FindPondId(id)
	if errFind == mongo.ErrNoDocuments {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "pond data not found"})
		return
	}
	if errFind != nil {
		utils.ErrorReq(c, errFind)
		return
	}

	var recycleColl = db.Db.Collection("deleted_ponds")
	_, errQuery := recycleColl.InsertOne(context.TODO(), pond)
	if errQuery != nil {
		utils.ErrorReq(c, errQuery)
		return
	}
	_, errDel := coll.DeleteOne(context.TODO(), bson.D{{"id", id}})
	if errDel != nil {
		utils.ErrorReq(c, errQuery)
		return
	}

	c.IndentedJSON(http.StatusOK,
		gin.H{"message": "data succesfully deleted"})
}

// GET all recycled farms
func GetRecycledFarms(c *gin.Context) {
	AddStats("GET /recycled/farms", c.ClientIP())
	// find all entries
	var coll = db.Db.Collection("deleted_farms")
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

// GET all recycled ponds
func GetRecycledPonds(c *gin.Context) {
	AddStats("GET /recycled/ponds", c.ClientIP())
	// find all entries
	var coll = db.Db.Collection("deleted_ponds")
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
