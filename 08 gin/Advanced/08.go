package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// Define a struct for the incoming data
type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=4,max=20"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Age      int    `json:"age" validate:"gte=18,lte=99"`
}

var validate = validator.New()

func main() {
	r := gin.Default()

	// Handle a POST request to create a user
	r.POST("/create-user", func(c *gin.Context) {
		// Bind the request body to the CreateUserRequest struct
		var request CreateUserRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate the request data
		if err := validate.Struct(request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// If validation passes, create the user (in a real application)
		// For this example, we'll just print the received data
		fmt.Printf("Creating user: %+v\n", request)
		c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
	})

	r.Run(":8080")
}




















package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// CustomValidator is a custom validator struct that embeds the default validator
type CustomValidator struct {
	validator *validator.Validate
}

// Validate is a method to perform custom validation
func (cv *CustomValidator) Validate(i interface{}) error {
	// Use the default validation
	if err := cv.validator.Struct(i); err != nil {
		return err
	}

	// Custom validation: Check if the username is "admin"
	if user, ok := i.(*CreateUserRequest); ok {
		if user.Username == "admin" {
			return validator.NewErr("username", "cannot be 'admin'")
		}
	}

	return nil
}

// CustomErrFieldRequired is a custom error message for the "required" validation tag
var CustomErrFieldRequired = "{0} is a required field"

// Define a struct for the incoming data
type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=4,max=20"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Age      int    `json:"age" validate:"gte=18,lte=99"`
}

var validate *validator.Validate

func main() {
	r := gin.Default()

	// Use the custom validator
	validate = validator.New()
	validate.RegisterValidation("customValidation", customValidation)

	r.Use(func(c *gin.Context) {
		c.Set("validator", validate)
		c.Next()
	})

	// Use the custom validator in Gin
	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		// Recover and handle panic, if any
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}))

	// Handle a POST request to create a user
	r.POST("/create-user", func(c *gin.Context) {
		// Bind the request body to the CreateUserRequest struct
		var request CreateUserRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate the request data using the custom validator
		if err := validate.Struct(request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// If validation passes, create the user (in a real application)
		// For this example, we'll just print the received data
		fmt.Printf("Creating user: %+v\n", request)
		c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
	})

	r.Run(":8080")
}

// customValidation is a custom validation function
func customValidation(fl validator.FieldLevel) bool {
	// Custom validation: Check if the field value is "custom"
	return fl.Field().String() != "custom"
}

