package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Load HTML templates from the "templates" directory.
	r.LoadHTMLGlob("templates/*.html")

	// Define a route handler to render the HTML template.
	r.GET("/", renderPage)

	r.Run(":8080")
}

func renderPage(c *gin.Context) {
	// Data to be passed to the template.
	data := gin.H{
		"Title":    "Welcome to My Website",
		"ItemList": []string{"Item 1", "Item 2", "Item 3"},
	}

	// Render the HTML template and pass data to it.
	c.HTML(http.StatusOK, "index.html", data)
}
