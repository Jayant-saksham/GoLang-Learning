package main

import (
	"socials/db"
	_ "socials/docs"
	"socials/routes"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func main() {
	// Initialize the database
	err := db.InitDB()
	if err != nil {
		panic(err)
	}

	// Set up the Gin router
	router := routes.SetupRouter()

	// Serve Swagger UI at /swagger
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start the server
	router.Run(":8080")
}

// url := 'http://0.0.0.0:8080/swagger/index.html'