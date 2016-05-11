package t411

import (
	"fmt"
	"strconv"
)

// Torrent represent a torrent as return by the t411 API
type Torrent struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Category       string `json:"category"`
	Rewritename    string `json:"rewritename"`
	Seeders        string `json:"seeders"`
	Leechers       string `json:"leechers"`
	Comments       string `json:"comments"`
	IsVerified     string `json:"isVerified"`
	Added          string `json:"added"`
	Size           string `json:"size"`
	TimesCompleted string `json:"times_completed"`
	Owner          string `json:"owner"`
	Categoryname   string `json:"categoryname"`
	Categoryimage  string `json:"categoryimage"`
	Username       string `json:"username"`
	Privacy        string `json:"privacy"`
}

// Torrents is an arry of torrent. This is used for sorting.
type Torrents []Torrent

func (t *Torrent) String() string {
	return fmt.Sprintf("%s - %s (%s)", t.ID, t.Name, t.Seeders)
}

// BySeeder implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded Torrents value.
type BySeeder struct{ Torrents }

// Less implements the sort.Interface
func (s BySeeder) Less(i, j int) bool {
	seederI, _ := strconv.Atoi(s.Torrents[i].Seeders)
	seederJ, _ := strconv.Atoi(s.Torrents[j].Seeders)
	return seederI < seederJ
}

// Len implements the sort.Interface
func (t Torrents) Len() int {
	return len(t)
}

// Swap implements the sort.Interface
func (t Torrents) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
