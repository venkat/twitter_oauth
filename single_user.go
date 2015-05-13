package main

import (
	"fmt"
	"github.com/mrjones/oauth"
	"io/ioutil"
	"log"
)

const (
	ConsumerKey       = "<YOUR CONSUMER KEY HERE>"
	ConsumerSecret    = "<YOUR CONSUMER SECRET HERE>"
	AccessToken       = "<YOUR ACCESS TOKEN HERE>"
	AccessTokenSecret = "<YOUR ACCESS TOKEN SECRET HERE>"
)

func main() {
	consumer := oauth.NewConsumer(ConsumerKey,
		ConsumerSecret,
		oauth.ServiceProvider{})
	//NOTE: remove this line or turn off Debug if you don't
	//want to see what the headers look like
	consumer.Debug(true)
	//Roll your own AccessToken struct
	accessToken := &oauth.AccessToken{Token: AccessToken,
		Secret: AccessTokenSecret}
	twitterEndPoint := "https://api.twitter.com/1.1/statuses/mentions_timeline.json"
	response, err := consumer.Get(twitterEndPoint, nil, accessToken)
	if err != nil {
		log.Fatal(err, response)
	}
	defer response.Body.Close()
	fmt.Println("Response:", response.StatusCode, response.Status)
	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(respBody))
}
