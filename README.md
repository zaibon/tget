# tget
[![Build Status](https://travis-ci.org/zaibon/tget.svg?branch=master)](https://travis-ci.org/zaibon/gtet) [![Go Report Card](https://goreportcard.com/badge/github.com/zaibon/tget)](https://goreportcard.com/report/github.com/zaibon/tget)

tget is a small CLI that helps you download the torrents of your favorite TV show in one command.

tget is greatly inspired by https://github.com/Hito01/t411-cli
## Install
```sh
go get -u github.com/zaibon/tget
```

## Build
To build from scratch simply use the Makefile provided. You just need to fill the `credentials.mk` file with the credentials of T411. This is used in the building process to retrieve some information about the T411 API and bundle it into the binary.
Then simply do a `make`

## Configuration
tget uses https://www.betaseries.com/ and http://www.t411.me/ respectively as source of metadata and torrent. Before being able to use tget you need to create a configuration file that old information from these two platforms.
```toml
[torrent]
output_dir="ssh://user@host:22/home/user/torrents"
language="vostfr"
[t411]
token="814529:132:9a0bfhu5e7f96f38ef21e0ecde11342f"
[betaseries]
login="user"
password="supersecret"
APIKey="8134c1956b70"
```
- torrent  
`language`: Set your prefered language.  You can choose between:
  - anglais
  - français
  - muet
  - multi-fr
  - multi-qb
  - québécois
  - vfstfr
  - vostfr
  - voasta

  `output_dir` : defined where to store the downloaded torrent. Format supported are:
  - local. `/home/user/torrents`
  - remote `ssh://user@host:22/home/user/torrents`  
remote try to use ssh-agent if loaded to access your ssh keys.


- t411  
`token` : API token from T411. you can automaticly get it using tget :
`t411-auth -u user -p password`

- betaseries  
`login` : optional, only needed if you want to mark torrent as retreived  
`password`: optional, only needed if you want to mark torrent as retreived  
`APIKey`: required, you can get one from https://www.betaseries.com/api/

## Usage
tget support download of a full show a season or an episode.

To download all the episode of all the season, just specify the title of the show.  
```
tget download --show 'silicon valey'
```
To download all episodes from a single season, specify the title and the season number
```
tget download --show 'silicon valey' --season 1
```
And to download only one episode, specify title, season and episode number.
```
tget download --show 'silicon valey' --season 1 --episode 5
```
