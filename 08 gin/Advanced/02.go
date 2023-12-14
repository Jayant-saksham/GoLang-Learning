package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

func main() {
	// Create a Gin router
	r := gin.Default()

	// Create a store for sessions (you can also use other storage options)
	store := cookie.NewStore([]byte("secret")) // "secret" is a secret key for encrypting session data

	// Use the sessions middleware with the store
	r.Use(sessions.Sessions("mySession", store))

	// Define a route that sets a session variable
	r.GET("/set-session", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Set("username", "JohnDoe")
		session.Save() // This saves the session data

		c.String(200, "Session data set!")
	})

	// Define a route that retrieves the session variable
	r.GET("/get-session", func(c *gin.Context) {
		session := sessions.Default(c)
		username := session.Get("username")

		if username == nil {
			c.String(200, "Username not found in session")
		} else {
			c.String(200, fmt.Sprintf("Username: %s", username))
		}
	})

	// Run the web server
	r.Run(":8080")
}















package main

import (
  "github.com/gin-contrib/sessions"
  "github.com/gin-contrib/sessions/cookie"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  store := cookie.NewStore([]byte("secret"))
  sessionNames := []string{"a", "b"}
  r.Use(sessions.SessionsMany(sessionNames, store))

  r.GET("/hello", func(c *gin.Context) {
    sessionA := sessions.DefaultMany(c, "a")
    sessionB := sessions.DefaultMany(c, "b")

    if sessionA.Get("hello") != "world!" {
      sessionA.Set("hello", "world!")
      sessionA.Save()
    }

    if sessionB.Get("hello") != "world?" {
      sessionB.Set("hello", "world?")
      sessionB.Save()
    }

    c.JSON(200, gin.H{
      "a": sessionA.Get("hello"),
      "b": sessionB.Get("hello"),
    })
  })
  r.Run(":8000")
}










package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// Initialize session store
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	// Dummy user data for demonstration (you should use a database)
	users := map[string]string{
		"user1": "password1",
		"user2": "password2",
	}

	// Home page
	r.GET("/", func(c *gin.Context) {
		session := sessions.Default(c)
		username := session.Get("username")

		if username != nil {
			c.HTML(http.StatusOK, "<h1>Welcome, "+username.(string)+"!</h1>")
		} else {
			c.HTML(http.StatusOK, "<h1>Welcome, Guest!</h1>")
		}
	})

	// Login page
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "<h1>Login Page</h1><form method='POST' action='/authenticate'><input type='text' name='username' placeholder='Username'><input type='password' name='password' placeholder='Password'><button type='submit'>Login</button></form>")
	})

	// Authenticate user
	r.POST("/authenticate", func(c *gin.Context) {
		session := sessions.Default(c)
		username := c.PostForm("username")
		password := c.PostForm("password")

		// Check if user exists and password is correct (in a real app, use a database)
		if storedPassword, found := users[username]; found && storedPassword == password {
			session.Set("username", username)
			session.Save()
			c.Redirect(http.StatusSeeOther, "/")
		} else {
			c.HTML(http.StatusOK, "<h1>Authentication Failed. Please try again.</h1>")
		}
	})

	// Logout
	r.GET("/logout", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Delete("username")
		session.Save()
		c.Redirect(http.StatusSeeOther, "/")
	})

	// Run the web server
	r.Run(":8000")
}
