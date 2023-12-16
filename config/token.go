package config

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

type TokenResp struct {
	AccessToken string `json:"access_token"`
}

func (c *Configuration) GetToken() {
	url := "https://aip.baidubce.com/oauth/2.0/token?client_id=" + c.Ak + "&client_secret=" + c.Sk + "&grant_type=client_credentials"
	payload := strings.NewReader(``)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		log.Fatalf("Error getting token %v", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error getting token %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error getting token %v", err)
	}
	var tokenResp TokenResp
	err = json.Unmarshal(body, &tokenResp)
	if err != nil {
		log.Fatalf("Error unmarshaling token %v", err)
	}
	c.Token = tokenResp.AccessToken
}
