package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testRedisService = &RedisService{}

func init() {
	testRedisService = InitializeCache()
}

func TestCacheInit(t *testing.T) {
	assert.True(t, testRedisService.redisClient != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "https://www.long-link-to-test.html"
	shortURL := "Jsz4k57oAX"
	SetURL(initialLink, shortURL)
	retrievedUrl := GetURL(shortURL)
	assert.Equal(t, initialLink, retrievedUrl)
}
