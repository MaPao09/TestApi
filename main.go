package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    // Create a Gin router instance
    r := gin.Default()

    // Define a route for the root path
    r.GET("/", func(c *gin.Context) {
        c.String(200, "คิคิ รักนะคะ")
    })

    // Start the HTTP server on port 8080
    r.Run(":8080")
}
