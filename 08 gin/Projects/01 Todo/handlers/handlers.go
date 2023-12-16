package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	// "01 Todo/models"
)

func postTodo(c * gin.Context) {
	
}

func listTodo(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html");
}

func deleteTodo(c *gin.Context) {
	//
}

func getTodo(c * gin.Context) {

}