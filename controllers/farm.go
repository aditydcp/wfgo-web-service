package controllers

import (
	"net/http"

	"aditydcp/wfgo-web-service/models"

	"github.com/gin-gonic/gin"
)

// to get one farm data with id
func (idb *InDB) GetFarmById(c *gin.Context) {
	var (
		farm   models.Farm
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&farm).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": farm,
		}
	}

	c.JSON(http.StatusOK, result)
}

// get all data of farms
func (idb *InDB) GetAllFarms(c *gin.Context) {
	var (
		farms  []models.Farm
		result gin.H
	)

	idb.DB.Find(&farms)
	if len(farms) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": farms,
			"count":  len(farms),
		}
	}

	c.JSON(http.StatusOK, result)
}

// create new data to the database
func (idb *InDB) CreateFarm(c *gin.Context) {
	var (
		farm   models.Farm
		result gin.H
	)
	id := c.PostForm("id")
	name := c.PostForm("name")
	farm.Id = id
	farm.Name = name
	idb.DB.Create(&farm)
	result = gin.H{
		"result": farm,
	}
	c.JSON(http.StatusOK, result)
}

// update farm data of given id
func (idb *InDB) UpdateFarm(c *gin.Context) {
	id := c.Query("id")
	name := c.PostForm("name")

	var (
		farm    models.Farm
		newFarm models.Farm
		result  gin.H
	)

	err := idb.DB.First(&farm, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}

	newFarm.Name = name

	err = idb.DB.Model(&farm).Updates(newFarm).Error
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

// delete farm data with given id
func (idb *InDB) DeleteFarm(c *gin.Context) {
	var (
		farm   models.Farm
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&farm, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	err = idb.DB.Delete(&farm).Error
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
