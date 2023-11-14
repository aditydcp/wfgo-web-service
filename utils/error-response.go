package utils

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func ErrorReq(c *gin.Context, err error) {
	c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
}