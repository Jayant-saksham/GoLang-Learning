package router

import (
	"Social/internal/app"
	"github.com/gin-gonic/gin"
)

func SetupRouter(a *app.Application) *gin.Engine {
	r := gin.Default()

	// User routes
	userRouter := r.Group("/users")
	{
		userRouter.POST("/", a.UserHandler.CreateUser)
		userRouter.GET("/:id", a.UserHandler.GetUserByID)
		userRouter.PUT("/:id", a.UserHandler.UpdateUser)
		userRouter.DELETE("/:id", a.UserHandler.DeleteUser)
	}

	// Post routes
	postRouter := r.Group("/posts")
	{
		postRouter.POST("/", a.PostHandler.CreatePost)
		postRouter.GET("/:id", a.PostHandler.GetPostByID)
		postRouter.PUT("/:id", a.PostHandler.UpdatePost)
		postRouter.DELETE("/:id", a.PostHandler.DeletePost)
	}

	// Comment routes
	commentRouter := r.Group("/comments")
	{
		commentRouter.POST("/", a.CommentHandler.CreateComment)
		commentRouter.GET("/:id", a.CommentHandler.GetCommentByID)
		commentRouter.PUT("/:id", a.CommentHandler.UpdateComment)
		commentRouter.DELETE("/:id", a.CommentHandler.DeleteComment)
	}

	return r
}
