package controllers

import (
	"net/http"

	"aditydcp/wfgo-web-service/models"

	"github.com/gin-gonic/gin"
)

// Get all farms data
func getFarms(c *gin.Context) {
	var farms []models.Farm
	models.DB.Find(&farms)

	c.JSON(http.StatusOK, gin.H{"result": farms})
}
