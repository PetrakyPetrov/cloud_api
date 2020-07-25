package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/petrakypetrov/cloud_clients_api/models"
)

// GetFolders ..
func GetFolders(c *gin.Context) {

	var folder models.Folder

	UserID := c.Param("user_id")
	ID, err := strconv.ParseInt(UserID, 10, 64)
	folder.UserID = ID

	result, err := folder.GetAllByUserID()
	if err != nil {
		fmt.Println(err)
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

// CreateFolders ..
func CreateFolders(c *gin.Context) {
	var folder models.Folder

	bytess, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server error",
			"data":    []string{},
		})
		return
	}

	if err := json.Unmarshal(bytess, &folder); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server error",
			"data":    []string{},
		})
		return
	}

	result, err := folder.Create()
	if err != nil {
		fmt.Println(err)
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

// DeleteFolders ..
func DeleteFolders(c *gin.Context) {
	var folder models.Folder

	UserID := c.Param("user_id")
	FolderID := c.Param("folder_id")

	FID, err := strconv.ParseInt(FolderID, 10, 64)
	if err == nil {
		fmt.Printf("%d of type %T", FID, FID)
	}

	UID, err := strconv.ParseInt(UserID, 10, 64)
	if err == nil {
		fmt.Printf("%d of type %T", FID, FID)
	}
	folder.UserID = UID
	folder.ID = FID

	result, err := folder.Delete()
	if err != nil {
		fmt.Println(err)
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
