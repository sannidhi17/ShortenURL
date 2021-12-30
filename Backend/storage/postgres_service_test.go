package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testPostgresService = &PostgresService{}

func init() {
	testPostgresService = InitDB()
}

func TestCacheInit(t *testing.T) {
	assert.True(t, testPostgresService.postgresdb != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "https://www.long-link-to-test.html"
	shortURL := "Jsz4k57oAX"
	AddURLToDB(initialLink, shortURL)
	retrievedUrl := GetURLFromDB(shortURL)
	assert.Equal(t, initialLink, retrievedUrl)
}
