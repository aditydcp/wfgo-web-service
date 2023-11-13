package controllers

import (
	"net/http"

	"aditydcp/wfgo-web-service/models"

	"github.com/gin-gonic/gin"
)

// to get one pond data with id
func (idb *InDB) GetPondById(c *gin.Context) {
	var (
		pond   models.Pond
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&pond).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": pond,
		}
	}

	c.JSON(http.StatusOK, result)
}

// get all data of ponds
func (idb *InDB) GetAllPonds(c *gin.Context) {
	var (
		ponds  []models.Pond
		result gin.H
	)

	idb.DB.Find(&ponds)
	if len(ponds) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": ponds,
			"count":  len(ponds),
		}
	}

	c.JSON(http.StatusOK, result)
}

// create new data to the database
func (idb *InDB) CreatePond(c *gin.Context) {
	var (
		pond   models.Pond
		result gin.H
	)
	id := c.PostForm("id")
	name := c.PostForm("name")
	pond.Id = id
	pond.Name = name
	idb.DB.Create(&pond)
	result = gin.H{
		"result": pond,
	}
	c.JSON(http.StatusOK, result)
}

// update pond data of given id
func (idb *InDB) UpdatePond(c *gin.Context) {
	id := c.Query("id")
	name := c.PostForm("name")

	var (
		pond    models.Pond
		newPond models.Pond
		result  gin.H
	)

	err := idb.DB.First(&pond, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}

	newPond.Name = name

	err = idb.DB.Model(&pond).Updates(newPond).Error
	if err != nil {
		result = gin.H{
			"result": "update failed",
		}
	} else {
		result = gin.H{
			"result": "successfully updated data",
		}
	}
	c.JSON(http.StatusOK, result)
}

// delete pond data with given id
func (idb *InDB) DeletePond(c *gin.Context) {
	var (
		pond   models.Pond
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&pond, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	err = idb.DB.Delete(&pond).Error
	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
	} else {
		result = gin.H{
			"result": "Data deleted successfully",
		}
	}

	c.JSON(http.StatusOK, result)
}
