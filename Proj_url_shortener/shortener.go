package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Shortener this is Sai Kumar URL",
		})
	})
	err := r.Run(":8081")
	if err != nil {
		panic(fmt.Sprintf("Failed to Start the web server-error: %v", err))
	}
}
