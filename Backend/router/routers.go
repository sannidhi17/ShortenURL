package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shortenurl/shorturl"
	"shortenurl/storage"
)

// Request model definition
type RequestModel struct {
	OriginalUrl string `json:"original_url" binding:"required"`
}

const (
	host = "http://localhost:8080/"
	nores = "No Result"
)

func ShortenURL(c *gin.Context) {
	var creationRequest RequestModel
	err := c.ShouldBindJSON(&creationRequest);
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	shortURL := shorturl.GenerateShortURL(creationRequest.OriginalUrl)
	originalUrl := storage.GetURLFromDB(shortURL)
	if (originalUrl == nores) {
		storage.AddURLToDB(creationRequest.OriginalUrl, shortURL)
	}

	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortURL,
	})
}

func RedirectToOriginalURL(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	originalUrl := storage.GetURLFromDB(shortUrl)
	if (originalUrl != nores) {
		c.Redirect(302, originalUrl)
	} else {
		c.JSON(400, "Invalid param")
	}
}