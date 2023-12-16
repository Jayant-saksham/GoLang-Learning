package routes

import (
	"socials/db"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	db.InitDB()

	api := r.Group("/api")

	// Use sessions middleware with cookie store
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	// Post routes
	postRoutes := api.Group("/posts")
	{
		postRoutes.GET("", GetPosts)
		postRoutes.POST("", CreatePost)
		postRoutes.GET("/:id", GetPost)
		postRoutes.PUT("/:id", UpdatePost)
		postRoutes.DELETE("/:id", DeletePost)
	}

	// Comment routes
	commentRoutes := api.Group("/comments")
	{
		commentRoutes.GET("", GetComments)
		commentRoutes.POST("", CreateComment)
		commentRoutes.GET("/:id", GetComment)
		commentRoutes.PUT("/:id", UpdateComment)
		commentRoutes.DELETE("/:id", DeleteComment)
	}

	// User routes
	userRoutes := api.Group("/users")
	{
		userRoutes.GET("", GetUsers)
		userRoutes.POST("", CreateUser)
		userRoutes.GET("/:id", GetUser)
		userRoutes.PUT("/:id", UpdateUser)
		userRoutes.DELETE("/:id", DeleteUser)
	}

	// Authentication route
	authRoutes := api.Group("/auth")
	{
		authRoutes.POST("/register", RegisterUser)
		authRoutes.POST("/login", LoginUser)
		authRoutes.POST("/logout", LogoutUser)
	}
	return r
}
