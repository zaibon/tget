package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"sync"

	"os"

	"github.com/zaibon/tget/api/betaseries"
	"github.com/zaibon/tget/api/t411"
)

func downloadShow(title string) error {
	cfg := loadConfig()
	t411API := t411.New(cfg.T411Token)
	bs := betaseries.New(cfg.BetaseriesToken)

	show, err := bs.Show(title)
	if err != nil {
		log.Printf("Error search information about the show: %v", err.Error())
		return err
	}

	wg := sync.WaitGroup{}
	for season, detail := range show.SeasonsDetails {
		for episode := 1; episode < detail.Episodes; episode++ {

			wg.Add(1)
			go func(title string, season, episode int) {
				defer wg.Done()

				name := fmt.Sprintf("%s_S%02dE%02d", title, season, episode)
				oldPath, err := t411API.DownloadTorrent(title, season, episode)
				if err != nil {
					log.Printf("Error downloading %s", name)
				}

				newPath := storePath(cfg.TorrentDirectory, name)
				if err = move(oldPath, newPath); err != nil {
					log.Printf("Error moving %s to %s : %v", oldPath, newPath, err)
				}

			}(title, season, episode)
		}
	}

	wg.Wait()

	return nil
}

func downloadSaison(title string, season int) error {
	cfg := loadConfig()
	t411API := t411.New(cfg.T411Token)
	bs := betaseries.New(cfg.BetaseriesToken)

	show, err := bs.Show(title)
	if err != nil {
		log.Printf("Error search information about the show: %v", err.Error())
		return err
	}

	log.Printf("%+v", show)

	detail := show.SeasonsDetails[season-1]
	wg := sync.WaitGroup{}

	for episode := 1; episode < detail.Episodes; episode++ {

		wg.Add(1)
		go func(title string, season, episode int) {
			defer wg.Done()

			name := fmt.Sprintf("%s_S%02dE%02d", title, season, episode)
			oldPath, err := t411API.DownloadTorrent(title, season, episode)
			if err != nil {
				log.Printf("Error downloading %s", name)
			}

			newPath := storePath(cfg.TorrentDirectory, name)
			if err = move(oldPath, newPath); err != nil {
				log.Printf("Error moving %s to %s : %v", oldPath, newPath, err)
			}

		}(title, season, episode)
	}

	wg.Wait()

	return nil

}

func downloadEpisode(title string, season, episode int) error {
	cfg := loadConfig()
	t411API := t411.New(cfg.T411Token)

	name := fmt.Sprintf("%s_S%02dE%02d", title, season, episode)
	oldPath, err := t411API.DownloadTorrent(title, season, episode)
	if err != nil {
		log.Printf("Error downloading %s", name)
	}

	newPath := storePath(cfg.TorrentDirectory, name)
	if err = move(oldPath, newPath); err != nil {
		log.Printf("Error moving %s to %s : %v", oldPath, newPath, err)
	}

	return nil
}

func storePath(dir, name string) string {
	path := filepath.Clean(filepath.Join(dir, name+".torrent"))
	path = strings.Replace(path, " ", "_", -1)
	return path
}

func move(oldPath, newPath string) error {
	data, err := ioutil.ReadFile(oldPath)
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile(newPath, data, 0660); err != nil {
		return err
	}
	return os.Remove(oldPath)
}
