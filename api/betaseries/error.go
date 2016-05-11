package betaseries

import "fmt"

// ErrAPI represents an error returned by the API
type errAPI struct {
	Errors []struct {
		Code int    `json:"code"`
		Text string `json:"text"`
	} `json:"errors"`
}

func (e errAPI) Error() string {
	out := ""
	for _, e := range e.Errors {
		out += fmt.Sprintf("%s\n", e.Text)
	}
	return out
}
