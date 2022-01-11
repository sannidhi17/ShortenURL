package main

import (
	"fmt"
	"shortenurl/router"
	"shortenurl/storage"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	GinRouter := gin.Default()
	GinRouter.Use(cors.Default())
	GinRouter.GET("/", getResponse)

	GinRouter.POST("/create-short-url", func(c *gin.Context) {
		router.ShortenURL(c)
	})

	GinRouter.GET("/:shortUrl", func(c *gin.Context) {
		router.RedirectToOriginalURL(c)
	})

	storage.InitDB()

	err := GinRouter.Run()
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	} else {
		fmt.Println("Server started!! :)")
	}

}

func getResponse(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello Go Shortener!",
	})
}
