package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/petrakypetrov/cloud_clients_api/models"
)

// GetFiles ..
func GetFiles(c *gin.Context) {

	var file models.File

	UserID := c.Param("user_id")
	file.UserID = UserID

	result, err := file.GetAllByUserID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server error",
			"data":    []string{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    result,
	})
}
