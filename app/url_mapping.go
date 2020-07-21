package app

import "github.com/petrakypetrov/cloud_clients_api/controllers"

func mapUrls() {

	router.GET("/ping", controllers.Ping)

	// group: v1
	v1 := router.Group("/api/v1")
	{
		v1.GET("/users/:user_id", controllers.GetUser)
		// v1.GET("/users", controllers.FindUser)
		v1.POST("/users", controllers.CreateUser)
	}
}
