package main

// This script make a request on t411 API to get the complete list of all the terms used for the research in the API.
// This create a file containing the mapping between season/episode number and the corresponding term used by t411.
// the mapping is then bundled into the source code of tget using go-bindata to be used during runtime.
// You need to specify your t411 login/password to be able to generate the token used by the API.

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"
)

type tmp struct {
	Type  string
	Mode  string
	Terms map[string]string
}

type intKV struct {
	Key   int `json:"key"`
	Value int `json:"value"`
}

type strKV struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

type dataMap struct {
	Seasons   []intKV `json:"seasons"`
	Episodes  []intKV `json:"episodes"`
	Languages []strKV `json:"languages"`
}

var login string
var password string

func init() {
	flag.StringVar(&login, "login", "", "t411 login")
	flag.StringVar(&password, "password", "", "t411 password")
}

func main() {
	flag.Parse()

	token, err := auth(login, password)
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest("GET", "https://api.t411.ch/terms/tree", nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Authorization", token)
	req.Header.Set("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	data := map[string]map[string]tmp{}
	defer resp.Body.Close()
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatalln(err)
	}

	out := dataMap{}
	extractEpisode(data, &out)
	extractSeason(data, &out)
	extractLanguage(data, &out)

	blob, err := json.MarshalIndent(out, "", "   ")
	if err != nil {
		log.Fatalln(err)
	}

	outputFilename := filepath.Join("scripts/mapping.json")
	if err = ioutil.WriteFile(outputFilename, blob, 0660); err != nil {
		log.Fatalln(err)
	}
    log.Println("mapping.json creared")
}

// Auth does a authentification request on T411 to get a new Token used to access the API
func auth(username, password string) (string, error) {
	u, err := url.Parse("https://api.t411.ch/auth")
	if err != nil {
		log.Println("Error creating auth url: ", err)
		return "", err
	}

	credentials := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{
		login,
		password,
	}

	buff := bytes.Buffer{}
	log.Println("Try authentification")
	if err = json.NewEncoder(&buff).Encode(credentials); err != nil {
		log.Println("Error encoding json auth request: ", err)
		return "", err
	}

	form := url.Values{}
	form.Set("username", credentials.Username)
	form.Set("password", credentials.Password)
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
		Error string `json:"error"`
	}{}

	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Println("Error decoding json auth response: ", err)
		return "", err
	}

	if data.Error != "" {
		return "", fmt.Errorf(data.Error)
	}

	if data.Token == "" {
		err := fmt.Errorf("Torken is empty.")
		log.Printf(err.Error())
		return "", err
	}

	return data.Token, nil
}

func extractEpisode(data map[string]map[string]tmp, out *dataMap) {
	for num, episode := range data["637"]["46"].Terms {
		strEPNr := strings.TrimLeft(episode, "Episode ")

		intEpNr, err := strconv.Atoi(strEPNr)
		if err != nil {
			continue
		}
		intNum, err := strconv.Atoi(num)
		if err != nil {
			continue
		}

		ep := intKV{
			Key:   intEpNr,
			Value: intNum,
		}

		out.Episodes = append(out.Episodes, ep)
	}
}

func extractSeason(data map[string]map[string]tmp, out *dataMap) {
	for num, season := range data["637"]["45"].Terms {
		strEPNr := strings.TrimLeft(season, "Saison ")

		intEpNr, err := strconv.Atoi(strEPNr)
		if err != nil {
			continue
		}
		intNum, err := strconv.Atoi(num)
		if err != nil {
			continue
		}

		ep := intKV{
			Key:   intEpNr,
			Value: intNum,
		}

		out.Seasons = append(out.Seasons, ep)
	}
}

func extractLanguage(data map[string]map[string]tmp, out *dataMap) {
	for num, language := range data["637"]["17"].Terms {
		intNum, err := strconv.Atoi(num)
		if err != nil {
			continue
		}

		ep := strKV{
			Key:   language,
			Value: intNum,
		}

		out.Languages = append(out.Languages, ep)
	}
}
