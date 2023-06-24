package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/imvahid/go-crud/initializers"
	"github.com/imvahid/go-crud/models"
)

func PostsCreate(c *gin.Context) {
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	var posts []models.Post

	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	id := c.Param("id")

	var post models.Post

	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get id of url
	id := c.Param("id")

	// Get the data of request body
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)

	// Find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get the id of the url
	id := c.Param("id")

	// Delete the posts
	initializers.DB.Delete(&models.Post{}, id)

	// Respond
	c.Status(200)
}
