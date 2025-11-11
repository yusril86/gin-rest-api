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

	// Get post with that id
	var post models.Post
	initializer.DB.First(&post, id)

	//Return it
	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostsUpdate(c *gin.Context) {
	// Get id from url
	id := c.Param("id")

	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	// Find the post we are updating
	var post models.Post
	initializer.DB.First(&post, id)

	// Update the post
	initializer.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	//Return it
	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostDelete(c *gin.Context) {

	// Get id from url
	id := c.Param("id")

	// Get post with that id
	var post models.Post
	initializer.DB.First(&post, id)

	// Delete the post
	initializer.DB.Delete(&post)

	//Return it
	c.JSON(200, gin.H{
		"message": "Post deleted successfully",
	})

}
