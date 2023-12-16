package routes

import (
	"net/http"

	"socials/db"
	"socials/models"
	"socials/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// RegisterUser handles user registration
func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password before storing it
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Store the hashed password in the database
	_, err = db.DB.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)",
		user.Username, user.Email, hashedPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// LoginUser handles user login
func LoginUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Query the database to check if the user credentials are valid
	var storedPassword string
	var storedUserID int
	err := db.DB.QueryRow("SELECT id, password FROM users WHERE username=?", user.Username).
		Scan(&storedUserID, &storedPassword)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Compare the stored password hash with the provided password
	err = utils.VerifyPassword(storedPassword, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Create a session and store the user ID
	session := sessions.Default(c)
	session.Set("userID", storedUserID)
	session.Save()

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

// LogoutUser handles user logout
func LogoutUser(c *gin.Context) {
	// Delete the session to log out the user
	session := sessions.Default(c)
	session.Delete("userID")
	session.Save()

	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
