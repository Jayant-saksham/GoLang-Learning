package routes

import (
	"net/http"
	"socials/db"
	"socials/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

//	@Summary		Get a list of comments
//	@Description	Get a list of comments
//	@Tags			Comments
//	@Produce		json
//	@Success		200	{array}	models.Comment
//	@Router			/api/comments [get]
func GetComments(c *gin.Context) {
	var comments []models.Comment

	rows, err := db.DB.Query("SELECT * FROM comments")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var comment models.Comment
		if err := rows.Scan(&comment.ID, &comment.Text, &comment.PostID, &comment.UserID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}

//	@Summary		Create a new comment
//	@Description	Create a new comment
//	@Tags			Comments
//	@Accept			json
//	@Produce		json
//	@Param			input	body		models.Comment	true	"Comment input"
//	@Success		201		{string}	string			"Comment created successfully"
//	@Router			/api/comments [post]
func CreateComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := db.DB.Exec("INSERT INTO comments (text, post_id, user_id) VALUES (?, ?, ?)",
		comment.Text, comment.PostID, comment.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	commentID, _ := result.LastInsertId()
	comment.ID = int(commentID)

	c.JSON(http.StatusCreated, comment)
}

//	@Summary		Get a specific comment by ID
//	@Description	Get a specific comment by ID
//	@Tags			Comments
//	@Produce		json
//	@Param			id	path		int	true	"Comment ID"
//	@Success		200	{object}	models.Comment
//	@Router			/api/comments/{id} [get]
func GetComment(c *gin.Context) {
	commentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	var comment models.Comment
	err = db.DB.QueryRow("SELECT * FROM comments WHERE id=?", commentID).
		Scan(&comment.ID, &comment.Text, &comment.PostID, &comment.UserID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	c.JSON(http.StatusOK, comment)
}

//	@Summary		Update an existing comment by ID
//	@Description	Update an existing comment by ID
//	@Tags			Comments
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"Comment ID"
//	@Param			input	body		models.Comment	true	"Comment input"
//	@Success		200		{string}	string			"Comment updated successfully"
//	@Router			/api/comments/{id} [put]
func UpdateComment(c *gin.Context) {
	commentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = db.DB.Exec("UPDATE comments SET text=?, post_id=?, user_id=? WHERE id=?",
		comment.Text, comment.PostID, comment.UserID, commentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment updated successfully"})
}

//	@Summary		Delete a comment by ID
//	@Description	Delete a comment by ID
//	@Tags			Comments
//	@Produce		json
//	@Param			id	path		int		true	"Comment ID"
//	@Success		200	{string}	string	"Comment deleted successfully"
//	@Router			/api/comments/{id} [delete]
func DeleteComment(c *gin.Context) {
	commentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	_, err = db.DB.Exec("DELETE FROM comments WHERE id=?", commentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}
