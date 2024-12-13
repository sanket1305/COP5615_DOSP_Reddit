package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"reddit_client_server/lozapi"
)

func main() {
	client := lozapi.NewClient(lozapi.BaseUrl, &http.Client{
		Timeout: 10 * time.Second,
	})

	response, err := client.GetMonsters()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Message: %s\n", response.Message)
}
