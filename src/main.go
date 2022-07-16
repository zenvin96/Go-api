package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func PostHomePage(c *gin.Context) {
	message := c.PostForm("message")
	c.JSON(200, gin.H{
		"message": message,
	})
}

func QueryString(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func main() {
	fmt.Println("Hello, World!")

	r := gin.Default()

	r.GET("/", HomePage)
	r.GET("/query", QueryString)
	r.POST("/PostHomePage", PostHomePage)
	r.GET("/nihao", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Nihao!"},
		)
	})

	r.Run("8090")
}
