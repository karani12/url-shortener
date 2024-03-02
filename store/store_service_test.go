package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testStorageService = &StorageService{}

func init() {
	testStorageService = InitializeStore()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStorageService.redisClient != nil)
}

func TestSaveUrlMapping(t *testing.T) {
	initialLink := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	userUUId := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	shortLink := "Jsz4k57oAX"

	err := SaveUrlMapping(shortLink, initialLink, userUUId)
	if err != nil {
		return
	}

	assert.True(t, RetrieveInitialUrl(shortLink) == initialLink)

}
