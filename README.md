# tget
[![Build Status](https://travis-ci.org/zaibon/tget.svg?branch=master)](https://travis-ci.org/zaibon/gtet) [![Go Report Card](https://goreportcard.com/badge/github.com/zaibon/tget)](https://goreportcard.com/report/github.com/zaibon/tget)

tget is a small CLI that helps you download the torrents of your favorite TV show in one command.

tget is greatly inspired by https://github.com/Hito01/t411-cli
## Install
```
go get -u github.com/zaibon/tget
```

## Configuration
tget uses https://www.betaseries.com/ and http://www.t411.me/ respectivly as source of metadata and torrent. Before beeing able to use tget you need to create a configuration file.  tget provide a command to help you create this config file
```
tget config --t411-username toto --t411-password supersecret --bs-token 2345345 --output /home/user/downloads/torrents
```
`--t411-username` 	t411 username  
`--t411-password` 	t411 password  
`--bs-token` 		betaseries API key. go to https://www.betaseries.com/api/ to get one  
`--output` 		directory where to save the downloaded torrents  


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
