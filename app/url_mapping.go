package app

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/petrakypetrov/cloud_clients_api/controllers"
	"github.com/petrakypetrov/cloud_clients_api/models"
)

// BaseAuthMiddleware ...
func BaseAuthMiddleware(c *gin.Context) {
	val, exist := c.Request.Header["Cs-Token"]

	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Error",
			"data":    []string{},
		})
		c.Abort()
		return
	}

	csToken := val[0]
	if len(strings.TrimSpace(csToken)) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Error",
			"data":    []string{},
		})
		c.Abort()
		return
	}

	var user models.User

	user.Token = csToken
	result, err := user.GetByToken()
	if err != nil || len(result) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"data":    result,
		})
		c.Abort()
		return
	}

}

func mapUrls() {

	router.GET("/ping", controllers.Ping)
	router.POST("/users/:user_id/files/upload", controllers.UploadFiles)
	// group: v1
	v1 := router.Group("/api/v1")
	{

		v1.POST("/users", controllers.CreateUser)
		v1.POST("/user/login", controllers.GetUserByEmailPassword)

		v1.Use(BaseAuthMiddleware)
		v1.GET("/users/:user_id/files", controllers.GetFiles)
		v1.GET("/users/:user_id", controllers.GetUser)
	}
}
