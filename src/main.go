package main

import (
	"go-api/src/controller"
	"go-api/src/database"
	"go-api/src/middlewares"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controller.GenerateToken)
		api.POST("/user/register", controller.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controller.Ping)
		}
	}
	return router
}

func main() {

	database.Connect("root:root@tcp(localhost:3306)/sohaidatabase?parseTime=true")
	database.Migrate()
	// Initialize Router
	router := initRouter()
	router.Run(":8080")

}
