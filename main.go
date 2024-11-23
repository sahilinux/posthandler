package main

import "github.com/gin-gonic/gin"

type User struct {
	Name  string `json:"name" binding:"required"` // Added 'binding' for validation
	Email string `json:"email" binding:"required,email"`
}

func main() {
	r := gin.Default()

	// Corrected route name to '/users'
	r.POST("/users", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			// Return validation error with detailed message
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, gin.H{
			"message": "User created successfully!",
			"user":    user,
		})
	})

	// GET route to retrieve users
	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This is the GET /users endpoint. Add functionality here.",
		})
	})
	r.Run(":8080") // Runs on port 8080
}
