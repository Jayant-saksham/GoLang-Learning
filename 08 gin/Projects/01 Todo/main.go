package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"handlers/handlers.go"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*html")
	r.GET("/", listTodo)
	r.GET("/:todoID", getTodo)
	r.POST("/:todoID", postTodo)
	r.DELETE("/:todoID", deleteTodo)
	r.Run(":8000")
}
