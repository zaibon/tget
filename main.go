package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"
	"github.com/naoina/toml"
	"github.com/zaibon/tget/api/t411"
)

type config struct {
	TorrentDirectory string
	T411             struct {
		Token string
	}
	Betaseries struct {
		Login    string
		Password string
		APIKey   string
	}
}

func loadConfig() (config config) {
	home := os.Getenv("HOME")
	f, err := os.Open(filepath.Join(home, ".tget.toml"))
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Please create configuration file at %s", f.Name())
			os.Exit(1)
		} else {
			log.Fatal(err)
		}
	}

	if err := toml.NewDecoder(f).Decode(&config); err != nil {
		log.Fatal(err)
	}
	return
}

func main() {

	app := cli.NewApp()
	app.Name = "autot411"

	var (
		title    string
		season   int
		episode  int
		t411Auth t411.AuthPair
		auto     = false
	)

	app.Commands = []cli.Command{
		{
			Name:  "t411-auth",
			Usage: "get t411 token",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "u,username",
					Usage:       "t411 username",
					Destination: &t411Auth.Username,
				},
				cli.StringFlag{
					Name:        "p, password",
					Usage:       "t411 password",
					Destination: &t411Auth.Password,
				},
			},
			Action: func(c *cli.Context) {

				if t411Auth.Username == "" || t411Auth.Password == "" {
					fmt.Println("Missing arguments")
					return
				}

				var err error
				cfg := config{}

				home := os.Getenv("HOME")
				path := filepath.Join(home, ".tget.toml")

				data, err := ioutil.ReadFile(path)
				if err == nil {
					if err = toml.Unmarshal(data, &cfg); err != nil {
						log.Fatal(err)
					}
				}

				cfg.T411.Token, err = t411.Auth(t411Auth)
				if err != nil {
					log.Fatal(err)
				}

				if data, err = toml.Marshal(cfg); err != nil {
					log.Fatal(err)
				}
				if err = ioutil.WriteFile(path, data, 0660); err != nil {
					log.Fatal(err)
				}
				log.Printf("config file updated (%s)", path)
			},
		},
		{
			Name:  "download",
			Usage: "download a show",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "show, t",
					Usage:       "Title of the show to dowload",
					Destination: &title,
				},
				cli.IntFlag{
					Name:        "season, s",
					Usage:       "Season number",
					Destination: &season,
				},
				cli.IntFlag{
					Name:        "episode, e",
					Usage:       "episode number",
					Destination: &episode,
				},
				cli.BoolFlag{
					Name:        "auto",
					Usage:       "automaticly download the torrent with more seeders. don't ask",
					Destination: &auto,
				},
			},
			Action: func(c *cli.Context) {

				if title == "" {
					fmt.Println("Please specify a title")
					return
				}

				if season == 0 && episode == 0 {
					downloadShow(title)
					return
				}

				if season != 0 && episode == 0 {
					downloadSaison(title, season)
					return
				}

				downloadEpisode(title, season, episode)
			},
		},
	}

	app.Run(os.Args)
}
