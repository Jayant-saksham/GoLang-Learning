package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default()

    // Define a dynamic route that accepts a username parameter.
    r.GET("/profile/:username", func(c *gin.Context) {
        // Get the username parameter from the URL.
        username := c.Param("username")

        // Simulate fetching user data from a database.
        // In a real application, you would perform a database query here.
        user := getUserByUsername(username)

        if user == nil {
            // User not found.
            c.String(http.StatusNotFound, "User not found")
            return
        }

        // Respond with the user's profile.
        c.JSON(http.StatusOK, gin.H{
            "username": user.Username,
            "email":    user.Email,
            // Add more user details as needed.
        })
    })

    r.Run(":8080")
}

// Simulated user data.
type User struct {
    Username string
    Email    string
}

// Simulated database function.
func getUserByUsername(username string) *User {
    // In a real application, you would fetch user data from a database.
    // For this example, we return predefined user data.
    users := map[string]*User{
        "john": {
            Username: "john",
            Email:    "john@example.com",
        },
        "jane": {
            Username: "jane",
            Email:    "jane@example.com",
        },
    }

    return users[username]
}
