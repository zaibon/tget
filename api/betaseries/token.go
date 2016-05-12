package betaseries

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

//token is struct return by betaseries when requesting a token
type token struct {
	User struct {
		ID        int    `json:"id"`
		Login     string `json:"login"`
		InAccount bool   `json:"in_account"`
	} `json:"user"`
	Token  string        `json:"token"`
	Hash   string        `json:"hash"`
	Errors []interface{} `json:"errors"`
}

func hashPasswd(passwd string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(passwd)))
}

func (b *BetaSeries) getToken() error {
	u, err := url.Parse(b.baseURL + "/members/auth")
	if err != nil {
		log.Fatal(err)
	}

	q := u.Query()
	q.Set("login", b.login)
	q.Set("password", hashPasswd(b.password))
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		log.Printf("Error creating request for %s: %v", u.String(), err.Error())
		return err
	}

	resp, err := b.do(req)
	if err != nil {
		log.Printf("Error getting token :%v", err)
		return err
	}
	if resp.StatusCode != http.StatusOK {
		apiErr := decodeErr(resp.Body)
		log.Fatalln(apiErr.Error())
		return apiErr
	}

	data := token{}

	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Printf("Error decoding token response :%v", err)
		return err
	}

	b.token = data.Token
	return nil
}
