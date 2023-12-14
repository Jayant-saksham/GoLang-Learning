package main // This says we're making a new program.

import ( // We're using some special tools to help us.
  "net/http" // We're using something to help us with the internet.
  "github.com/gin-gonic/gin" // We're using a special helper called "gin" from the internet.
)

func main() { // This is where our program starts.

  r := gin.Default() // We create something called "r" to help us with making a website.

  r.GET("/ping", func(c *gin.Context) { // We're saying that if someone goes to our website and types "/ping" in their web browser...
    c.JSON(http.StatusOK, gin.H{ // ...we'll send them a message.
      "message": "pong", // The message will say "pong."
    })
  })

  r.Run() // We tell our website to start running on the internet.
}

















type MyData struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

func MyHandler(c *gin.Context) {
    var b MyData

    // Bind the JSON data from the request to the MyData struct.
    if err := c.Bind(&b); err != nil {
        // Handle the error (e.g., invalid JSON data).
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Now, the 'b' variable contains the data from the request.
    // You can use 'b.Name' and 'b.Email' to access the values.
    c.JSON(http.StatusOK, b)
}
