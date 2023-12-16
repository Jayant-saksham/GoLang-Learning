package routes

import (
	"net/http"
	"strconv"

	"socials/db"
	"socials/models"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

//	@Summary		Get a list of posts
//	@Description	Get a list of posts
//	@Tags			Posts
//	@Produce		json
//	@Success		200	{array}	models.Post
//	@Router			/api/posts [get]
func GetPosts(c *gin.Context) {
	var posts []models.Post

	rows, err := db.DB.Query("SELECT * FROM posts")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.UserID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}

//	@Summary		Create a new post
//	@Description	Create a new post
//	@Tags			Posts
//	@Accept			json
//	@Produce		json
//	@Param			input	body		models.Post	true	"Post input"
//	@Success		201		{string}	string		"Post created successfully"
//	@Router			/api/posts [post]
func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.DB.Exec("INSERT INTO posts (title, content, user_id) VALUES (?, ?, ?)",
		post.Title, post.Content, post.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Post created successfully"})
}

//	@Summary		Get a specific post by ID
//	@Description	Get a specific post by ID
//	@Tags			Posts
//	@Produce		json
//	@Param			id	path		int	true	"Post ID"
//	@Success		200	{object}	models.Post
//	@Router			/api/posts/{id} [get]
func GetPost(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	var post models.Post
	err = db.DB.QueryRow("SELECT * FROM posts WHERE id=?", postID).
		Scan(&post.ID, &post.Title, &post.Content, &post.UserID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

//	@Summary		Update an existing post by ID
//	@Description	Update an existing post by ID
//	@Tags			Posts
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int			true	"Post ID"
//	@Param			input	body		models.Post	true	"Post input"
//	@Success		200		{string}	string		"Post updated successfully"
//	@Router			/api/posts/{id} [put]
func UpdatePost(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = db.DB.Exec("UPDATE posts SET title=?, content=?, user_id=? WHERE id=?",
		post.Title, post.Content, post.UserID, postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully"})
}

//	@Summary		Delete a post by ID
//	@Description	Delete a post by ID
//	@Tags			Posts
//	@Produce		json
//	@Param			id	path		int		true	"Post ID"
//	@Success		200	{string}	string	"Post deleted successfully"
//	@Router			/api/posts/{id} [delete]
func DeletePost(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	_, err = db.DB.Exec("DELETE FROM posts WHERE id=?", postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
