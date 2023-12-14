type UserInput struct {
    Name  string `json:"name" binding:"required"`
    Email string `json:"email" binding:"required,email"`
    Age   int    `json:"age" binding:"required,numeric"`
}












package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// UserInput represents the expected JSON input structure.
type UserInput struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Age   int    `json:"age" binding:"required,numeric"`
}

var validate = validator.New()

func main() {
	r := gin.Default()

	// Define a route for creating a user.
	r.POST("/create-user", createUser)

	r.Run(":8080")
}

func createUser(c *gin.Context) {
	var userInput UserInput

	// Bind and validate the JSON request data.
	if err := c.ShouldBindJSON(&userInput); err != nil {
		// Validation failed, return a 400 Bad Request response with error details.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Data is valid, you can use userInput.Name, userInput.Email, userInput.Age.
	// Process and store the data as needed.

	// For this example, we'll just return a success message.
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}


