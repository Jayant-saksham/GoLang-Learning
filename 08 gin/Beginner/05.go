package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default()

    // Define a route that accepts a "name" query parameter.
    r.GET("/greet", func(c *gin.Context) {
        // Get the value of the "name" query parameter from the URL.
        name := c.Query("name")

        if name == "" {
            // Handle the case when the "name" parameter is missing.
            c.String(http.StatusBadRequest, "Please provide a name")
            return
        }

        // Respond with a personalized greeting.
        greeting := "Hello, " + name + "!"
        c.String(http.StatusOK, greeting)
    })

    r.Run(":8080")
}
