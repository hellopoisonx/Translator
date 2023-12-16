package src

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"translator/config"
)

type Result struct {
	Destination string `json:"dst"`
	Source      string `json:"src"`
}

type Resp struct {
	Result struct {
		TransResult []Result `json:"trans_result"`
		From        string   `json:"from"`
		To          string   `json:"to"`
	} `json:"result"`
	LogID int64 `json:"log_id"`
}

const BaseURL = "https://aip.baidubce.com/rpc/2.0/mt/texttrans/v1"

func Translate(args []string) {
	conf := config.NewConfiguration()
	conf.ParseConfiguration(args)
	req, err := json.Marshal(&conf.Req)
	if err != nil {
		log.Fatalf("Error marshaling the config %v", err)
	}
	reqURL := BaseURL + "?access_token=" + conf.Token
	reqU, err := http.NewRequest("POST", reqURL, bytes.NewBuffer(req))
	if err != nil {
		log.Fatalf("Error posting request %v", err)
	}
	reqU.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(reqU)
	if err != nil {
		log.Fatalf("Error posting request %v", err)
	}
	defer resp.Body.Close()
	respContent, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error parsing response %v", err)
	}
	var respS Resp
	err = json.Unmarshal(respContent, &respS)
	if err != nil {
		log.Fatalf("Error ummarshaling response %v", err)
	}
	fmt.Println(respS.Result.TransResult)
}
