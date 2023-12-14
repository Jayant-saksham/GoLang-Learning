// Middleware functions
func Middleware1(c *gin.Context) {
    // Your middleware code
    c.Next()
}

func Middleware2(c *gin.Context) {
    // Your middleware code
    c.Next()
}

// Route with multiple middleware and a handler
r.GET("/admin", Middleware1, Middleware2, func(c *gin.Context) {
    // Your code to handle the GET request and generate a response
})












package main

import (
    "github.com/gin-gonic/gin"
    "time"
    "fmt"
)

func LoggerMiddleware() gin.HandlerFunc {
    // This is our custom middleware function.
    return func(c *gin.Context) {
        // Before handling the request, let's log some information.
        method := c.Request.Method     // Get the HTTP method (GET, POST, etc.).
        url := c.Request.RequestURI    // Get the requested URL.
        timestamp := time.Now().Format("2006-01-02 15:04:05")

        // Log the request details.
        fmt.Printf("[%s] %s %s\n", timestamp, method, url)

        // Continue processing the request.
        c.Next() // This line tells Gin to proceed to the next middleware or handler.
    }
}

func main() {
    r := gin.Default() // Create a Gin router.

    // Use the LoggerMiddleware for all routes.
    r.Use(LoggerMiddleware())

    // Define a route that responds to a GET request.
    r.GET("/hello", func(c *gin.Context) {
        // When a GET request is made to /hello, this function is executed.
        c.String(200, "Hello, World!")
    })

    // Start the Gin web server.
    r.Run(":8080")
}
