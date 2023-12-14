package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Todo structure for our TODO items
type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var todos []Todo

func main() {
	r := gin.Default()

	// Define routes
	r.POST("/todos", createTodo)
	r.GET("/todos", getTodos)
	r.GET("/todos/:id", getTodo)
	r.PUT("/todos/:id", updateTodo)
	r.DELETE("/todos/:id", deleteTodo)

	// Run the web server
	r.Run(":8000")
}

// createTodo handles the creation of a new TODO item
func createTodo(c *gin.Context) {
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo.ID = len(todos) + 1
	todos = append(todos, todo)

	c.JSON(http.StatusCreated, todo)
}

// getTodos returns a list of all TODO items
func getTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

// getTodo returns a single TODO item by its ID
func getTodo(c *gin.Context) {
	id := c.Param("id")
	for _, todo := range todos {
		if id == string(todo.ID) {
			c.JSON(http.StatusOK, todo)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "TODO item not found"})
}

// updateTodo updates a TODO item by its ID
func updateTodo(c *gin.Context) {
	id := c.Param("id")
	var updatedTodo Todo

	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, todo := range todos {
		if id == string(todo.ID) {
			todos[i] = updatedTodo
			c.JSON(http.StatusOK, updatedTodo)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "TODO item not found"})
}

// deleteTodo deletes a TODO item by its ID
func deleteTodo(c *gin.Context) {
	id := c.Param("id")

	for i, todo := range todos {
		if id == string(todo.ID) {
			todos = append(todos[:i], todos[i+1:]...)
			c.JSON(http.StatusNoContent, nil)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "TODO item not found"})
}
