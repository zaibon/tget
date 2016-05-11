package betaseries

//Show represents a Show as returned by the betaserie API
// currently, unused attributes are commented
type Show struct {
	ID int `json:"id"`
	// ThetvdbID      int    `json:"thetvdb_id"`
	// ImdbID         string `json:"imdb_id"`
	Title string `json:"title"`
	// Description    string `json:"description"`
	Seasons        string `json:"seasons"`
	SeasonsDetails []struct {
		Number   int `json:"number"`
		Episodes int `json:"episodes"`
	} `json:"seasons_details"`
	Episodes string `json:"episodes"`
	// 	Followers  string   `json:"followers"`
	// 	Comments   string   `json:"comments"`
	// 	Similars   string   `json:"similars"`
	// 	Characters string   `json:"characters"`
	// 	Creation   string   `json:"creation"`
	// 	Genres     []string `json:"genres"`
	// 	Length     string   `json:"length"`
	// 	Network    string   `json:"network"`
	// 	Rating     string   `json:"rating"`
	// 	Status     string   `json:"status"`
	// 	Language   string   `json:"language"`
	// 	Notes      struct {
	// 		Total string `json:"total"`
	// 		Mean  int    `json:"mean"`
	// 		User  int    `json:"user"`
	// 	} `json:"notes"`
	// 	InAccount bool `json:"in_account"`
	// 	Images    struct {
	// 		Poster string `json:"poster"`
	// 	} `json:"images"`
	// 	Aliases []interface{} `json:"aliases"`
	// 	User    struct {
	// 		Archived  bool        `json:"archived"`
	// 		Favorited bool        `json:"favorited"`
	// 		Remaining int         `json:"remaining"`
	// 		Status    int         `json:"status"`
	// 		Last      string      `json:"last"`
	// 		Tags      interface{} `json:"tags"`
	// 	} `json:"user"`
	// 	ResourceURL string `json:"resource_url"`
}
