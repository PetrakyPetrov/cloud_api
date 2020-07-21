package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/petrakypetrov/cloud_clients_api/models"
)

// CreateUser ..
func CreateUser(c *gin.Context) {
	var user models.User

	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server error",
			"data":    []string{},
		})
		return
	}

	if err := json.Unmarshal(bytes, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server error",
			"data":    []string{},
		})
		return
	}

	if len(strings.TrimSpace(user.Email)) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email cannot be ampty",
			"data":    []string{},
		})
		return
	}

	if len(strings.TrimSpace(user.Password)) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Password cannot be ampty",
			"data":    []string{},
		})
		return
	}

	result, err := user.Create()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"data":    []string{},
		})
		return
	}

	type UserToken struct {
		Value string `json:"value"`
	}

	var ut UserToken
	ut.Value = result.Token
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    []UserToken{ut},
	})
}

// GetUserByPassword ..
func GetUserByPassword(c *gin.Context) {

	var user models.User

	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server error",
			"data":    []string{},
		})
		return
	}

	if err := json.Unmarshal(bytes, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server error",
			"data":    []string{},
		})
		return
	}

	if len(strings.TrimSpace(user.Email)) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email cannot be ampty",
			"data":    []string{},
		})
		return
	}

	if len(strings.TrimSpace(user.Password)) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Password cannot be ampty",
			"data":    []string{},
		})
		return
	}

	result, err := user.GetByPass()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"data":    []string{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    result,
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
