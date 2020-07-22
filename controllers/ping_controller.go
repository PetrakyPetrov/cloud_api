package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping ...
func Ping(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "OK",
		"data":    []string{},
	})
}
