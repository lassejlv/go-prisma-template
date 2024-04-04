package main

import (
	"goprismatemp/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Load environment variables
	err := godotenv.Load()

	// Check if there was an error loading the .env file
	if err != nil {
		panic("Error loading .env file")
	}

	// Prisma setup
	client := db.NewClient()

	// Connect to the database
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}

	// Close the connection when the main function exits
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	// ctx := context.Background()

	// Server stuff
	r := gin.Default()

	r.GET("/posts", func(c *gin.Context) {
		posts, err := client.Post.FindMany().Exec(c)

		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		} else {
			c.JSON(200, gin.H{
				"posts": posts,
			})
		}
	})

	r.POST("/posts", func(c *gin.Context) {
		// Ponst title, and published
		title := c.PostForm("title")
		published := c.PostForm("published")

		// Check if title and published are empty
		if (title == "") || (published == "") {
			c.JSON(400, gin.H{
				"error": "Title and published fields are required",
			})
			return
		}

		// Create a new post
		post, err := client.Post.CreateOne(
			db.Post.Title.Set(title),
			db.Post.Published.Set(published == "true"),
		).Exec(c)

		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		} else {
			// Return the created post
			c.JSON(201, gin.H{
				"post": post,
			})
		}

	})

	// Handle 404 requests
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"error": "Route was not found on the server :(",
		})
	})

	r.Run()
}
