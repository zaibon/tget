package t411

import (
	"fmt"
	"log"
	"net/url"
)

// T411 search API is quite strange to use. see https://api.t411.ch/
// they use 'terms' to allow search by category.
// In this case we are only interested in category Season and episode number.
// Season and episode number also have specific ID. init method creates the mapping

var (
	catSeasonID  = 45
	catEpisodeID = 46
	seasonNbrID  = map[int]int{}
	episodeNbrID = map[int]int{}
)

func init() {
	for i := 0; i < 30; i++ {
		seasonNbrID[i+1] = 968 + i
	}
	for i := 0; i < 60; i++ {
		episodeNbrID[i+1] = 937 + i
	}
}

type searchReq struct {
	Title   string
	Season  int
	Episode int
}

// URL returns the url of the search request
func (r searchReq) URL() string {
	u, err := url.Parse(fmt.Sprintf("%s/torrents/search/%s", t411BaseURL, r.Title))
	if err != nil {
		log.Fatalf("Error during construction of t411 search URL: %v", err)
	}
	q := u.Query()
	if r.Season > 0 {
		q.Add(fmt.Sprintf("term[%d][]", catSeasonID), fmt.Sprintf("%d", seasonNbrID[r.Season]))
	}
	if r.Episode > 0 {
		q.Add(fmt.Sprintf("term[%d][]", catEpisodeID), fmt.Sprintf("%d", episodeNbrID[r.Episode]))
	}
	u.RawQuery = q.Encode()

	return u.String()
}
