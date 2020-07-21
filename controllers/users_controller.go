package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/petrakypetrov/cloud_clients_api/models"
)

// CreateUser ..
func CreateUser(c *gin.Context) {
	var user models.User

	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		// TODO: Handle error
		return
	}

	if err := json.Unmarshal(bytes, &user); err != nil {
		// TODO: Handle error
		fmt.Println("ima neshto")
		return
	}
	arr := []string{}
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    arr,
	})
}

// GetUser ..
func GetUser(c *gin.Context) {
	var user models.User

	result, err := user.Get()
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    result,
	})
}

// // FindUser ..
// func FindUser(c *gin.Context) {
// 	c.JSON(http.StatusNotImplemented, gin.H{
// 		"message": "Nema implementaciq bako",
// 	})
// }
