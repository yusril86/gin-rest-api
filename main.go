package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yusril86/gin-rest-api/initializer"
)

func init() {
	initializer.LoadEnvVariablesss()
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "testing",
			"status":  "success",
		})
	})
	router.Run() // listen and serve on 0.0.0.0:8080)
}
