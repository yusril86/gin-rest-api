package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "iasa",
		})
	})
	router.Run() // listen and serve on 0.0.0.0:8080)
}
