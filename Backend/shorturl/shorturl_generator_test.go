package shorturl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShortLinkGenerator(t *testing.T) {
	initialLink_1 := "https://www.test-link1/123.html"
	shortLink_1 := GenerateShortURL(initialLink_1)

	initialLink_2 := "https://www.test-link2-long-version/123.html"
	shortLink_2 := GenerateShortURL(initialLink_2)

	initialLink_3 := "https://www.test-link3/123/home-robots/hello-robots-stretch-mobile-manipulator"
	shortLink_3 := GenerateShortURL(initialLink_3)

	assert.Equal(t, shortLink_1, "CNgT3ymR")
	assert.Equal(t, shortLink_2, "b8TpJqC7")
	assert.Equal(t, shortLink_3, "UyB9KUs5")
}