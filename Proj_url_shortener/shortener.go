package main

import (
	"fmt"

	"github.com/bskbala410@gmail.com/go-url-shortener/Proj_url_shortener/handler"
	"github.com/bskbala410@gmail.com/go-url-shortener/Proj_url_shortener/store"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Shortener this is Sai Kumar URL",
		})
	})
	r.POST("/create-short-url",func(c *gin.Context){
		handler.CreateShortUrl(c)
	})
	r.GET("/:shortUrl",func(c *gin.Context){
		handler.HandlerShortUrlRedirect(c)

	})
	//Note that store Initializationhappens here
	store.Intializestore()
	
	err := r.Run(":8081")
	if err != nil {
		panic(fmt.Sprintf("Failed to Start the web server-error: %v", err))
	}
}
