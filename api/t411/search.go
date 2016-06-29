//go:generate go-bindata -nometadata -prefix ../../scripts/ -pkg t411 ../../scripts/mapping.json

package t411

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

// T411 search API is quite strange to use. see https://api.t411.ch/
// they use 'terms' to allow search by category.
// In this case we are only interested in category Season and episode number.
// Season and episode number also have specific ID. init method creates the mapping

var (
	catSeasonID   = 45
	catEpisodeID  = 46
	catLanguageID = 51
)

type t411Mapping struct {
	Seasons   map[int]int
	Episodes  map[int]int
	Languages map[string]int
	// Qualities map[string]int
}

var mapping t411Mapping

func init() {
	mapping = t411Mapping{
		Seasons:   make(map[int]int),
		Episodes:  make(map[int]int),
		Languages: make(map[string]int),
	}

	dataMap := struct {
		Seasons []struct {
			Key   int `json:"key"`
			Value int `json:"value"`
		} `json:"seasons"`
		Episodes []struct {
			Key   int `json:"key"`
			Value int `json:"value"`
		} `json:"episodes"`
		Languages []struct {
			Key   string `json:"key"`
			Value int    `json:"value"`
		} `json:"languages"`
	}{}

	data, err := Asset("mapping.json")
	if err != nil {
		log.Fatalln("Error loading t411 asset:", err)
	}

	if err = json.Unmarshal(data, &dataMap); err != nil {
		log.Fatalln("Error Unmarshaling t411 mapping", err)
	}
	for _, kv := range dataMap.Seasons {
		mapping.Seasons[kv.Key] = kv.Value
	}
	for _, kv := range dataMap.Episodes {
		mapping.Episodes[kv.Key] = kv.Value
	}
	for _, kv := range dataMap.Languages {
		mapping.Languages[kv.Key] = kv.Value
	}
}

type searchReq struct {
	Title    string
	Season   int
	Episode  int
	Language string
}

// URL returns the url of the search request
func (r searchReq) URL() string {
	u, err := url.Parse(fmt.Sprintf("%s/torrents/search/%s", t411BaseURL, r.Title))
	if err != nil {
		log.Fatalf("Error during construction of t411 search URL: %v", err)
	}
	q := u.Query()
	if r.Season > 0 {
		q.Add(fmt.Sprintf("term[%d][]", catSeasonID), fmt.Sprintf("%d", mapping.Seasons[r.Season]))
	}
	if r.Episode > 0 {
		q.Add(fmt.Sprintf("term[%d][]", catEpisodeID), fmt.Sprintf("%d", mapping.Episodes[r.Episode]))
	}
	if ID, ok := mapping.Languages[r.Language]; ok {
		q.Add(fmt.Sprintf("term[%d][]", catLanguageID), fmt.Sprintf("%d", ID))

	}
	u.RawQuery = q.Encode()

	return u.String()
}
