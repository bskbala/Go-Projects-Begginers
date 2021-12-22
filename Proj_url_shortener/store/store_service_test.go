import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testStoreService = &StorageService{}

func int() {
	testStoreService = Intializestore()

}
func TestStoreInit(t *testing.T) {
	asser.True(t, testStoreService.redisClient != nil)

}

func TestInsertionAndRetrival(t *testing.T) {
	intialLink := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	userUUID := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	shortUrl := "Jsz4k57oAX"

	//Persist data mapping
	SaveUrlMapping(shortUrl, InitialLink, userUUID)

	//Retrieve Initial URL
	retrieveUrl := RetrieveIntialUrl(shortUrl)

	assert.Equal(t, InitialLink, retrieveUrl)
}