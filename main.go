package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"
	"github.com/naoina/toml"
	"github.com/zaibon/tget/api/t411"
)

type config struct {
	BetaseriesToken  string
	T411Token        string
	TorrentDirectory string
}

func loadConfig() (config config) {
	home := os.Getenv("HOME")
	f, err := os.Open(filepath.Join(home, ".autot411.toml"))
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Please run autot411 config to configure the app.")
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
		title      string
		season     int
		episode    int
		t411Auth   t411.AuthPair
		bsToken    string
		torrentDir string
		auto       = false
	)

	app.Commands = []cli.Command{
		{
			Name:  "config",
			Usage: "configure API access to betaseries and t411",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "t411-username",
					Usage:       "t411 username",
					Destination: &t411Auth.Username,
				},
				cli.StringFlag{
					Name:        "t411-password",
					Usage:       "t411 password",
					Destination: &t411Auth.Password,
				},
				cli.StringFlag{
					Name:        "bs-token",
					Usage:       "betaseries API key. go to https://www.betaseries.com/api/ to get one",
					Destination: &bsToken,
				},
				cli.StringFlag{
					Name:        "output",
					Usage:       "directory where to save the downloaded torrents",
					Destination: &torrentDir,
				},
			},
			Action: func(c *cli.Context) {

				if bsToken == "" || t411Auth.Username == "" || t411Auth.Password == "" || torrentDir == "" {
					fmt.Println("Missing arguments")
					return
				}

				if _, err := os.Stat(torrentDir); err != nil {
					if os.IsNotExist(err) {
						os.MkdirAll(torrentDir, 0775)
					}
				}

				var err error
				cfg := config{
					BetaseriesToken:  bsToken,
					TorrentDirectory: torrentDir,
				}

				cfg.T411Token, err = t411.Auth(t411Auth)
				if err != nil {
					log.Fatal(err)
				}

				home := os.Getenv("HOME")
				path := filepath.Join(home, ".autot411.toml")
				fmt.Println("Write config file to ", path)
				f, err := os.OpenFile(filepath.Join(home, ".autot411.toml"), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0660)
				if err != nil {
					log.Fatal(err)
				}
				if err := toml.NewEncoder(f).Encode(cfg); err != nil {
					log.Fatal(err)
				}
				fmt.Println("Configuration done.")
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
