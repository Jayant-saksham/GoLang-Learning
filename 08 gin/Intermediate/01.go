package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default()

    // Create a group of routes with a common URL prefix.
    blogGroup := r.Group("/blog")

    // Middleware for the entire blog group.
    blogGroup.Use(authMiddleware)

    // Routes within the blog group.
    blogGroup.GET("/posts", listBlogPosts)
    blogGroup.GET("/post/:id", getBlogPost)
    blogGroup.POST("/post", createBlogPost)

    // Other non-blog-related routes.
    r.GET("/", homePage)
    r.GET("/store", storePage)

    r.Run(":8080")
}

func authMiddleware(c *gin.Context) {
    // Simulated authentication logic.
    // In a real application, you would handle user authentication here.
    // For this example, we'll just set a flag in the context.
    c.Set("authenticated", true)
    c.Next()
}

func listBlogPosts(c *gin.Context) {
    // Check if the user is authenticated.
    authenticated, _ := c.Get("authenticated")
    if authenticated == true {
        c.String(http.StatusOK, "List of blog posts")
    } else {
        c.String(http.StatusUnauthorized, "Unauthorized. Please log in.")
    }
}

func getBlogPost(c *gin.Context) {
    id := c.Param("id")
    c.String(http.StatusOK, "Blog post with ID "+id)
}

func createBlogPost(c *gin.Context) {
    // Check if the user is authenticated.
    authenticated, _ := c.Get("authenticated")
    if authenticated == true {
        c.String(http.StatusOK, "Blog post created successfully")
    } else {
        c.String(http.StatusUnauthorized, "Unauthorized. Please log in.")
    }
}

func homePage(c *gin.Context) {
    c.String(http.StatusOK, "Welcome to the home page")
}

func storePage(c *gin.Context) {
    c.String(http.StatusOK, "Welcome to the store")
}













package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default()

    // Create a group for product-related routes.
    productsGroup := r.Group("/products")
    productsGroup.GET("/", listProducts)
    productsGroup.GET("/:id", getProduct)

    // Create a group for order-related routes.
    ordersGroup := r.Group("/orders")
    ordersGroup.Use(authMiddleware) // Apply authentication middleware to the order group.
    ordersGroup.GET("/", listOrders)
    ordersGroup.POST("/", createOrder)

    // Create a group for user profile-related routes.
    userProfileGroup := r.Group("/profile")
    userProfileGroup.GET("/:username", getUserProfile)

    r.Run(":8080")
}

func authMiddleware(c *gin.Context) {
    // Simulated authentication logic.
    // In a real application, you would handle user authentication here.
    // For this example, we'll just set a flag in the context.
    c.Set("authenticated", true)
    c.Next()
}

func listProducts(c *gin.Context) {
    c.String(http.StatusOK, "List of products")
}

func getProduct(c *gin.Context) {
    id := c.Param("id")
    c.String(http.StatusOK, "Product with ID "+id)
}

func listOrders(c *gin.Context) {
    // Check if the user is authenticated.
    authenticated, _ := c.Get("authenticated")
    if authenticated == true {
        c.String(http.StatusOK, "List of orders")
    } else {
        c.String(http.StatusUnauthorized, "Unauthorized. Please log in.")
    }
}

func createOrder(c *gin.Context) {
    // Check if the user is authenticated.
    authenticated, _ := c.Get("authenticated")
    if authenticated == true {
        c.String(http.StatusOK, "Order created successfully")
    } else {
        c.String(http.StatusUnauthorized, "Unauthorized. Please log in.")
    }
}

func getUserProfile(c *gin.Context) {
    username := c.Param("username")
    c.String(http.StatusOK, "User profile for "+username)
}
