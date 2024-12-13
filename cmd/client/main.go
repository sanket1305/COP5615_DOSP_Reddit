package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"reddit_client_server/lozapi"
)

func main() {
	// create instance of client to call different APIs
	client := lozapi.NewClient(lozapi.BaseUrl, &http.Client{
		Timeout: 10 * time.Second,
	})

	// register user
	response, err := client.RegisterUser()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", response.Message)

	// create subreddit
	response, err = client.CreateSubreddit("USA")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", response.Message)
}
