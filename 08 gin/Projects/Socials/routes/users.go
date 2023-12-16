package routes

import (
	"net/http"
	"socials/db"
	"socials/models"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)
//	@Summary		Get a list of users
//	@Description	Get a list of users
//	@Tags			Users
//	@Produce		json
//	@Success		200	{array}	models.User
//	@Router			/api/users [get]
func GetUsers(c *gin.Context) {
	var users []models.User

	rows, err := db.DB.Query("SELECT * FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

//	@Summary		Create a new user
//	@Description	Create a new user
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			input	body		models.User	true	"User input"
//	@Success		201		{string}	string		"User created successfully"
//	@Router			/api/users [post]
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := db.DB.Exec("INSERT INTO users (username, email) VALUES (?, ?)",
		user.Username, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userID, _ := result.LastInsertId()
	user.ID = int(userID)

	c.JSON(http.StatusCreated, user)
}

//	@Summary		Get a specific user by ID
//	@Description	Get a specific user by ID
//	@Tags			Users
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	models.User
//	@Router			/api/users/{id} [get]
func GetUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User
	err = db.DB.QueryRow("SELECT * FROM users WHERE id=?", userID).
		Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

//	@Summary		Update an existing user by ID
//	@Description	Update an existing user by ID
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int			true	"User ID"
//	@Param			input	body		models.User	true	"User input"
//	@Success		200		{string}	string		"User updated successfully"
//	@Router			/api/users/{id} [put]
func UpdateUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = db.DB.Exec("UPDATE users SET username=?, email=? WHERE id=?",
		user.Username, user.Email, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

//	@Summary		Delete a user by ID
//	@Description	Delete a user by ID
//	@Tags			Users
//	@Produce		json
//	@Param			id	path		int		true	"User ID"
//	@Success		200	{string}	string	"User deleted successfully"
//	@Router			/api/users/{id} [delete]
func DeleteUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	_, err = db.DB.Exec("DELETE FROM users WHERE id=?", userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
