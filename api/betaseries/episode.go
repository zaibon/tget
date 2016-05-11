package betaseries

// Episode represents an episode as return by the betaseries API
type Episode struct {
	ID        int         `json:"id"`
	ThetvdbID int         `json:"thetvdb_id"`
	YoutubeID interface{} `json:"youtube_id"`
	Title     string      `json:"title"`
	Season    int         `json:"season"`
	Episode   int         `json:"episode"`
	Show      struct {
		ID        int    `json:"id"`
		ThetvdbID int    `json:"thetvdb_id"`
		Title     string `json:"title"`
	} `json:"show"`
	Code        string `json:"code"`
	Global      int    `json:"global"`
	Special     int    `json:"special"`
	Description string `json:"description"`
	Date        string `json:"date"`
	User        struct {
		Seen       bool `json:"seen"`
		Downloaded bool `json:"downloaded"`
	} `json:"user"`
	Comments  string        `json:"comments"`
	Subtitles []interface{} `json:"subtitles"`
}
