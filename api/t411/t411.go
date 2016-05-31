package t411

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sort"
)

var (
	errEmptyToken = errors.New("token empty")
	errEOF        = errors.New("No more torrents to find")
)

const t411BaseURL = "https://api.t411.ch"

// AuthPair is the couple username password required by t411 authentification API
type AuthPair struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// T411 structure represents the T411 API
type T411 struct {
	baseURL   string
	token     string
	AuthPair  AuthPair
	outputDir string
}

// New creates a new T411 object using token
func New(token string) *T411 {
	t411 := &T411{
		baseURL:   t411BaseURL,
		outputDir: "/tmp",
		token:     token,
	}
	return t411
}

func (t *T411) do(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", t.token)
	req.Header.Set("Accept", "application/json")

	return http.DefaultClient.Do(req)
}

// Auth does a authentification request on T411 to get a new Token used to access the API
func Auth(AuthPair AuthPair) (string, error) {
	u, err := url.Parse(t411BaseURL + "/auth")
	if err != nil {
		log.Println("Error creating auth url: ", err)
		return "", err
	}

	buff := bytes.Buffer{}
	log.Println("Try authentification")
	if err = json.NewEncoder(&buff).Encode(AuthPair); err != nil {
		log.Println("Error encoding json auth request: ", err)
		return "", err
	}

	form := url.Values{}
	form.Set("username", AuthPair.Username)
	form.Set("password", AuthPair.Password)
	resp, err := http.PostForm(u.String(), form)
	if err != nil {
		log.Println("Error post auth request: ", err)
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Authentification failure, statuscode :%d", resp.StatusCode)
	}

	data := struct {
		Token string `json:"token"`
	}{}

	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Println("Error decoding json auth response: ", err)
		return "", err
	}

	if data.Token == "" {
		log.Printf("Torken is empty.")
		return "", errEmptyToken
	}

	return data.Token, nil
}

// DownloadTorrent search the torrent corresponding to the title,
// season and episode number, download it and return the location of the file.
func (t *T411) DownloadTorrent(title string, season, episode int, language string) (string, error) {
	req := searchReq{
		Title:    title,
		Season:   season,
		Episode:  episode,
		Language: language,
	}

	torrents, err := t.search(req)
	if err != nil {
		log.Printf("Error search for torrent: %v", err.Error())
		return "", err
	}

	if len(torrents) < 1 {
		return "", fmt.Errorf("torrent not found, %sS%02dE%02d", title, season, episode)
	}

	sort.Sort(BySeeder{torrents})

	r, err := t.download(torrents[0].ID)
	if err != nil {
		return "", err
	}
	defer r.Close()

	tmpfile, err := ioutil.TempFile("", fmt.Sprintf("%sS%02dE%02d", title, season, episode))
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer tmpfile.Close()

	if _, err = io.Copy(tmpfile, r); err != nil {
		log.Println(err)
		return "", err
	}

	return tmpfile.Name(), nil
}

// DownloadTorrents search the torrent specified by the SearchReq and download the torrent file locally
// func (t *T411) DownloadTorrents(req SearchReq, auto bool) error {
//
// 	toDownloads := make([]string, 10)
// 	searchChoose := func(req SearchReq) error {
// 		torrents, err := t.search(req)
// 		if err != nil {
// 			return err
// 		}
//
// 		// no more torrents to find
// 		if len(torrents) <= 0 {
// 			return errEOF
// 		}
//
// 		if auto {
// 			toDownloads = append(toDownloads, torrents[0].ID)
// 		} else {
// 			id, err := chooseTorrent(torrents)
// 			if err != nil {
// 				return err
// 			}
// 			toDownloads = append(toDownloads, id)
// 		}
// 		return nil
// 	}
//
// 	var err error
//
// 	// no season specify, download all seasons
// 	if req.Season == 0 {
// 		// TODO: ask betaseries for number of season. this would slow down interaction.
// 		// need to do this async
// 		for i := 1; err != errEOF; i++ {
// 			// TODO: same, ask betaseries for number of episode
// 			for j := 1; err != errEOF; j++ {
// 				req.Season = i
// 				req.Episode = j
// 				err = searchChoose(req)
// 				if err != nil && err != errEOF {
// 					log.Printf("Error searching for torrent %s S%dE%d", req.Title, req.Season, req.Episode)
// 					return err
// 				}
// 			}
// 		}
// 		return nil
// 	}
//
// 	// season specify, but not episodes. download all episodes of the season
// 	if req.Season > 0 && req.Episode == 0 {
// 		for j := 0; err != errEOF; j++ {
// 			req.Episode = j
// 			err := searchChoose(req)
// 			if err != nil && err != errEOF {
// 				log.Printf("Error searching for torrent %s S%dE%d", req.Title, req.Season, req.Episode)
// 				return err
// 			}
// 		}
// 		return nil
// 	}
//
// 	// all specified, just download what is requested
// 	if err := searchChoose(req); err != nil {
// 		log.Printf("Error searching for torrent %s S%dE%d", req.Title, req.Season, req.Episode)
// 		return err
// 	}
//
// 	fmt.Println(toDownloads)
// 	t.download(toDownloads)
// 	return nil
// }

// func (t *T411) downloadTorrents(req SearchReq, auto bool) error {
// 	torrents, err := t.search(req)
// 	if err != nil {
// 		return err
// 	}
//
// 	sort.Sort(sort.Reverse(BySeeder{torrents}))
//
// 	if len(torrents) > 1 {
// 		toDownloads := make([]string, 0, len(torrents))
// 		if auto {
// 			toDownloads = append(toDownloads, torrents[0].ID)
// 		} else {
// 			id, err := chooseTorrent(torrents)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			toDownloads = append(toDownloads, id)
// 		}
// 		t.download(toDownloads)
// 	}
// 	return nil
// }

func (t *T411) download(ID string) (io.ReadCloser, error) {

	u, err := url.Parse(fmt.Sprintf("%s/torrents/download/%s", t.baseURL, ID))
	if err != nil {
		log.Println("Error parsing url: ", err)
		return nil, err
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Println("Error creating downlaod request: ", err)
		return nil, err
	}

	resp, err := t.do(req)
	if err != nil {
		log.Println("Error executing download request: ", err)
		return nil, err
	}

	return resp.Body, err
}

func (t *T411) search(searchReq searchReq) ([]Torrent, error) {

	req, err := http.NewRequest("GET", searchReq.URL(), nil)
	if err != nil {
		log.Printf("Error creating request to %s: %v", searchReq.URL(), err)
		return nil, err
	}

	resp, err := t.do(req)
	if err != nil {
		log.Printf("Error executing request to %s: %v", searchReq.URL(), err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatal("bad status code", resp.StatusCode)
	}

	data := struct {
		Torrents []Torrent `json:"torrents"`
	}{}

	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Torrents, nil
}

//
// func chooseTorrent(torrents []Torrent) (string, error) {
// 	if len(torrents) < 1 {
// 		return "", fmt.Errorf("no torrents to choose from")
// 	}
//
// 	sort.Sort(sort.Reverse(BySeeder{torrents}))
//
// 	for i, torrent := range torrents {
// 		fmt.Printf("%d %s\n", i+1, torrent.String())
// 	}
// 	fmt.Printf("Which torrent do you want ? (1-%d) : \n", len(torrents))
// 	var index int
// 	fmt.Scanf("%d", &index)
//
// 	return torrents[index-1].ID, nil
// }
