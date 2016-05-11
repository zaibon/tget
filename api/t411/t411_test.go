package t411

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {
	username, password := os.Getenv("T411_USERNAME"), os.Getenv("T411_PASSWORD")
	// only do this test if we have the credentil set
	if username == "" || password == "" {
		return
	}

	token, err := Auth(AuthPair{Username: username, Password: password})
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	assert.NotEqual(t, "", token)
}

func TestReqURL(t *testing.T) {
	req := searchReq{
		Title:   "vikings",
		Season:  1,
		Episode: 1,
	}

	assert.Equal(t, "vikings", req.Title)
	assert.Equal(t, 1, req.Season)
	assert.Equal(t, 1, req.Episode)

	expected := "https://api.t411.ch/torrents/search/vikings?term%5B45%5D%5B%5D=968&term%5B46%5D%5B%5D=937"
	assert.Equal(t, expected, req.URL())

	req = searchReq{
		Title:   "vikings",
		Season:  3,
		Episode: 5,
	}
	expected = "https://api.t411.ch/torrents/search/vikings?term%5B45%5D%5B%5D=970&term%5B46%5D%5B%5D=941"
	assert.Equal(t, expected, req.URL())
}
