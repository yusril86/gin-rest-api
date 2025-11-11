package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yusril86/gin-rest-api/controllers"
	"github.com/yusril86/gin-rest-api/initializer"
)

func init() {
	initializer.LoadEnvVariablesss()
	initializer.ConnectToDB()
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "testing",
			"status":  "success",
		})
	})

	router.POST("/posts", controllers.PostsCreate)
	router.GET("/posts", controllers.PostsIndex)
	router.GET("/posts/:id", controllers.PostsShow)
	router.PUT("/posts/:id", controllers.PostsUpdate)
	router.DELETE("/posts/:id", controllers.PostDelete)

	router.Run() // listen and serve on 0.0.0.0:8080)
}
