package config

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Configuration struct {
	Token string `json:"token"`
	Ak    string `json:"ak"`
	Sk    string `json:"sk"`
	Req
}
type Req struct {
	SrcLang string `json:"from"`
	DstLang string `json:"to"`
	Query   string `json:"q"`
	TermIds string `json:"termIds"`
}

func NewConfiguration() *Configuration {
	r := Req{
		SrcLang: "auto",
		DstLang: "en",
		Query:   "",
		TermIds: "",
	}
	return &Configuration{
		Token: "",
		Ak:    "",
		Sk:    "",
		Req:   r,
	}
}

func (c *Configuration) ParseConfiguration(args []string) {
	// get configuration from arguments
	args = args[1:]
	// fmt.Println(args)
	c.Query = args[0]
	configPath := "/etc/translator/config.json"
	for i, v := range args {
		if (v == "-c" || v == "--config") && i+1 < len(args) {
			configPath = args[i+1]
		} else if (v == "-s" || v == "--src") && i+1 < len(args) {
			c.SrcLang = args[i+1]
		} else if (v == "-d" || v == "--dst") && i+1 < len(args) {
			// fmt.Println(v, args[i+1])
			c.DstLang = args[i+1]
		} else if (v == "-t" || v == "--term") && i+1 < len(args) {
			c.TermIds = args[i+1]
		} else if v == "--token" && i+1 < len(args) {
			c.Token = args[i+1]
		}
	}
	// get configuration from file
	f, err := os.Open(configPath)
	// if err.Error() == os.ErrExist.Error() {
	// 	log.Printf("Configuration file doesn't exist: %s. The Default configuration will be used", configPath)
	// }
	if err != nil {
		log.Fatalf("Error opening the configuration file %v", err)
	}
	defer f.Close()
	conf, err := io.ReadAll(f)
	if err != nil {
		log.Fatalf("Error reading the configuration file %v", err)
	}
	err = json.Unmarshal(conf, c)
	if err != nil {
		log.Fatalf("Error ummarshaling configuration %v", err)
	}
	c.GetToken()
	// fmt.Println(c)
}
