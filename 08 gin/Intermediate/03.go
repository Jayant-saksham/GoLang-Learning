package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// Custom middleware to check if the user is authenticated.
func authMiddleware(c *gin.Context) {
    // Check if the user is authenticated. In a real app, you'd have your authentication logic here.
    isAuthenticated := true // Simulated for this example.

    if !isAuthenticated {
        // If the user is not authenticated, respond with a 401 Unauthorized status.
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized. Please log in."})
        c.Abort() // Abort further processing of the request.
        return
    }

    // If the user is authenticated, continue to the next middleware or route handler.
    c.Next()
}

func main() {
    r := gin.Default()

    // Apply the authMiddleware to specific routes or groups of routes.
    r.GET("/private", authMiddleware, privateRoute)
    r.GET("/public", publicRoute)

    r.Run(":8080")
}

func privateRoute(c *gin.Context) {
    c.String(http.StatusOK, "This is a private route.")
}

func publicRoute(c *gin.Context) {
    c.String(http.StatusOK, "This is a public route.")
}
