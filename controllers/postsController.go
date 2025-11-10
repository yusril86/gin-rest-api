package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yusril86/gin-rest-api/initializer"
	"github.com/yusril86/gin-rest-api/models"
)

func PostsCreate(c *gin.Context) {
	// Get data from request body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	// Create post model
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializer.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return

	}

	c.JSON(200, gin.H{
		"message": "Posts Create",
	})

}

func PostsIndex(c *gin.Context) {
	// Get all records
	var posts []models.Post
	initializer.DB.Find(&posts)

	//Return it
	c.JSON(200, gin.H{
		"posts": posts,
	})

}

func PostsShow(c *gin.Context) {
	// Get id from url
	id := c.Param("id")

	// Get records
	var post models.Post
	initializer.DB.First(&post, id)

	//Return it
	c.JSON(200, gin.H{
		"post": post,
	})

}
