package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"

	"github.com/jaisingh/go-etrade/etrade"
	"github.com/jaisingh/oauth"
)

func main() {

	/*
		config := etrade.OauthConfig{
			ConsumerKey:    oauth_consumer_key,
			ConsumerSecret: oauth_consumer_secret,
			AccessToken: oauth.AccessToken{
				Token:  "4nokhL4VUiVc0S86WVCz3pf6eG2qi0ew=",
				Secret: "V2c3if+Gw0MNobva0tanVv7WzNFXlcgY=",
			},
		}

		data, _ := json.MarshalIndent(config, "", "  ")
		ioutil.WriteFile("config.json", data, 0644)
	*/

	rawConfig, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	config := etrade.OauthConfig{}

	if err := json.Unmarshal(rawConfig, &config); err != nil {
		log.Fatal(err)
	}

	c := etrade.NewOauthClient(config, "http://localhost:8080")

	// Verify token
	response, err := c.Get(etrade.URL_ACCOUNTLIST, map[string]string{})
	response.Body.Close()
	if err != nil {
		accessToken := &oauth.AccessToken{}
		accessToken, err = c.RefreshToken(&c.Config.AccessToken)
		if err != nil {
			log.Println(err)
			requestToken, url, err := c.GetRequestTokenAndUrl("oob")
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("(1) Go to: " + url)
			fmt.Println("(2) Grant access, you should get back a verification code.")
			fmt.Println("(3) Enter that verification code here: ")

			cmd := exec.Command("open", url)
			cmd.Start()
			verificationCode := ""

			fmt.Scanln(&verificationCode)

			accessToken, err = c.AuthorizeToken(requestToken, verificationCode)
			if err != nil {
				log.Fatal(err)
			}
		}
		c.Config.AccessToken = *accessToken

		log.Printf("Writing out new config file")

		data, _ := json.MarshalIndent(c.Config, "", "  ")
		if err := ioutil.WriteFile("config.json", data, 0644); err != nil {
			log.Fatal(err)
		}
	}

	al, err := c.GetQuote("MSFT")
	if err != nil {
		log.Println(err)
	}

	log.Printf("%v\n", al)

}
