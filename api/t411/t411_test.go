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
	type output struct {
		Title   string
		Season  int
		Episode int
		URL     string
	}
	tt := []struct {
		input  searchReq
		expect output
	}{
		{
			input: searchReq{
				Title:   "vikings",
				Season:  1,
				Episode: 1,
			},
			expect: output{
				Title:   "vikings",
				Season:  1,
				Episode: 1,
				URL:     "https://api.t411.ch/torrents/search/vikings?term%5B45%5D%5B%5D=968&term%5B46%5D%5B%5D=937",
			},
		},
		{
			input: searchReq{
				Title:   "vikings",
				Season:  3,
				Episode: 11,
			},
			expect: output{
				Title:   "vikings",
				Season:  3,
				Episode: 11,
				URL:     "https://api.t411.ch/torrents/search/vikings?term%5B45%5D%5B%5D=970&term%5B46%5D%5B%5D=948",
			},
		},
	}

	for _, test := range tt {
		assert.Equal(t, test.expect.Title, test.input.Title)
		assert.Equal(t, test.expect.Episode, test.input.Episode)
		assert.Equal(t, test.expect.Season, test.input.Season)
		assert.Equal(t, test.expect.URL, test.input.URL())
	}
}
