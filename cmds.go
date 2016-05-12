package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/url"
	"strings"
	"sync"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"

	"os"

	"github.com/tmc/scp"
	"github.com/zaibon/tget/api/betaseries"
	"github.com/zaibon/tget/api/t411"
)

var (
	t411API   *t411.T411
	bsAPI     *betaseries.BetaSeries
	outputDir string
)

func init() {
	cfg := loadConfig()

	t411API = t411.New(cfg.T411.Token)
	bsAPI = betaseries.New(cfg.Betaseries.APIKey, cfg.Betaseries.Login, cfg.Betaseries.Password)
	outputDir = cfg.TorrentDirectory
}

func downloadShow(title string) error {

	show, err := bsAPI.Show(title)
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
				download(title, season, episode, outputDir)
			}(title, season, episode)
		}
	}

	wg.Wait()

	return nil
}

func downloadSaison(title string, season int) error {

	show, err := bsAPI.Show(title)
	if err != nil {
		log.Printf("Error search information about the show: %v", err.Error())
		return err
	}

	detail := show.SeasonsDetails[season-1]
	wg := sync.WaitGroup{}

	for episode := 1; episode < detail.Episodes; episode++ {

		wg.Add(1)
		go func(title string, season, episode int) {
			defer wg.Done()
			download(title, season, episode, outputDir)
		}(title, season, episode)
	}

	wg.Wait()

	return nil

}

func downloadEpisode(title string, season, episode int) error {
	return download(title, season, episode, outputDir)
}

func download(title string, season int, episode int, destDir string) error {
	name := fmt.Sprintf("%s_S%02dE%02d", title, season, episode)
	log.Printf("download %s", name)

	oldPath, err := t411API.DownloadTorrent(title, season, episode)
	if err != nil {
		log.Printf("Error downloading %s: %v", name, err)
		return err
	}

	newPath := storePath(destDir, name)
	if err = move(oldPath, newPath); err != nil {
		log.Printf("Error moving %s to %s : %v", oldPath, newPath, err)
		return err
	}

	if err := bsAPI.MarkDownloaded(title, season, episode); err != nil {
		log.Printf("Could not mark episode as downloaded :%v", err)
	}

	log.Printf("%s save at %s", name, newPath)

	return nil
}

func storePath(dir, name string) string {
	path := dir + "/" + name + ".torrent"
	path = strings.Replace(path, " ", "_", -1)
	return path
}

func move(oldPath, newPath string) error {
	defer os.Remove(oldPath)

	u, err := url.Parse(newPath)
	if err != nil {
		return err
	}

	if u.Scheme == "ssh" {
		return sshMove(oldPath, newPath)
	}

	return localMove(oldPath, newPath)
}

func localMove(oldPath, newPath string) error {
	data, err := ioutil.ReadFile(oldPath)
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile(newPath, data, 0660); err != nil {
		return err
	}
	return os.Remove(oldPath)
}

func sshMove(oldPath, newPath string) error {
	u, err := url.Parse(newPath)
	if err != nil {
		return err
	}

	// try to found way to authentify against the server
	authMethods := []ssh.AuthMethod{}
	if passwd, isSet := u.User.Password(); isSet {
		authMethods = append(authMethods, ssh.Password(passwd))
	}

	agent, err := getAgent()
	if err == nil {
		authMethods = append(authMethods, ssh.PublicKeysCallback(agent.Signers))
	}

	client, err := ssh.Dial("tcp", u.Host, &ssh.ClientConfig{
		User: u.User.Username(),
		Auth: authMethods,
	})
	if err != nil {
		return fmt.Errorf("Failed to dial: %s", err)
	}

	session, err := client.NewSession()
	if err != nil {
		return fmt.Errorf("Failed to create session: %s", err.Error())
	}

	err = scp.CopyPath(oldPath, u.Path, session)
	if err != nil {
		return fmt.Errorf("Failed writing file: %s", err)
	}

	return nil
}

func getAgent() (agent.Agent, error) {
	agentConn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
	return agent.NewClient(agentConn), err
}
