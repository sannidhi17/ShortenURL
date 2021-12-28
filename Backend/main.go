package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Printf("Hello ShortenURL\n")
	router := gin.Default()
	router.GET("/", getResponse)
	err := router.Run()
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	} else {
		fmt.Println("Server started!! :)")
	}
}

func getResponse(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hey Go Shortener!",
	})
}
