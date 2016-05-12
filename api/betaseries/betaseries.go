package betaseries

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var (
	errShowNotFound = errors.New("Show not found")
)

const (
	betaseriesBaseURL = "https://api.betaseries.com"
	version           = "2.4"
)

// BetaSeries is the main object that give acces to the BetaSeries API
type BetaSeries struct {
	baseURL  string
	version  string
	apiKey   string
	token    string
	login    string
	password string
}

// New creates a new BetaSeries object.
func New(APIKey string, login, password string) *BetaSeries {
	bs := &BetaSeries{
		baseURL:  betaseriesBaseURL,
		version:  version,
		apiKey:   APIKey,
		login:    login,
		password: password,
	}

	if login != "" && password != "" {
		bs.getToken()
	}
	return bs
}

func (b *BetaSeries) do(req *http.Request) (*http.Response, error) {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-BetaSeries-Version", b.version)
	req.Header.Set("X-BetaSeries-Key", b.apiKey)
	if b.token != "" {
		req.Header.Set("X-BetaSeries-Token", b.token)
	}

	return http.DefaultClient.Do(req)
}

// Show returns the complete information about a show
func (b *BetaSeries) Show(title string) (*Show, error) {
	u, err := url.Parse(b.baseURL + "/shows/search")
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	q.Set("title", strings.ToLower(title))
	q.Set("order", "popularity")
	q.Set("nbpp", "1")
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Printf("Error creating request for %s: %v", u.String(), err.Error())
		return nil, err
	}

	resp, err := b.do(req)
	if err != nil {
		log.Printf("Error getting showID of %s: %v", title, err.Error())
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		apiErr := decodeErr(resp.Body)
		log.Println(apiErr.Error())
		return nil, apiErr
	}

	data := struct {
		Shows []Show `json:"shows"`
	}{}

	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Printf("Error decoding showID for '%s' :%v", title, err)
		return nil, err
	}

	if len(data.Shows) < 1 {
		log.Println("shows empty")
		return nil, errShowNotFound
	}

	return &data.Shows[0], nil
}

// Episode returns the information about an episode
func (b *BetaSeries) Episode(title string, season, episode int) (*Episode, error) {
	u, err := url.Parse(b.baseURL + "/shows/episodes")
	if err != nil {
		log.Fatal(err)
	}

	show, err := b.Show(title)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("id", fmt.Sprintf("%d", show.ID))
	q.Set("season", fmt.Sprintf("%d", season))
	q.Set("episode", fmt.Sprintf("%d", episode))
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Printf("Error creating request for %s: %v", u.String(), err.Error())
		return nil, err
	}

	resp, err := b.do(req)
	if err != nil {
		log.Printf("Error executing request on %s: %v", u.String(), err.Error())
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		apiErr := decodeErr(resp.Body)
		log.Println(apiErr.Error())
		return nil, apiErr
	}

	data := struct {
		Episodes []*Episode
	}{}

	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Printf("Error decoding Episode list for %s %d %d : %v", title, season, episode, err)
		return nil, err
	}

	return data.Episodes[0], nil
}

// MarkDownloaded mark an episode as download
func (b *BetaSeries) MarkDownloaded(title string, season, episode int) error {
	ep, err := b.Episode(title, season, episode)
	if err != nil {
		return err
	}

	u, err := url.Parse(b.baseURL + "/episodes/downloaded")
	if err != nil {
		log.Fatal(err)
	}

	q := u.Query()
	q.Set("id", fmt.Sprintf("%d", ep.ID))
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		log.Printf("Error marking episode as downloaded: %v", err.Error())
		return err
	}

	resp, err := b.do(req)
	if err != nil {
		log.Printf("Error marking episode as downloaded: %v", err.Error())
		return err
	}
	if resp.StatusCode != http.StatusOK {
		apiErr := decodeErr(resp.Body)
		log.Fatalf("MarkDownloaded bad response code: %v", apiErr.Error())
		return apiErr
	}

	return nil
}

func decodeErr(r io.Reader) (err errAPI) {
	if jsonerr := json.NewDecoder(r).Decode(&err); jsonerr != nil {
		log.Fatalf("Error decoding API error : %v", jsonerr)
	}
	return
}
