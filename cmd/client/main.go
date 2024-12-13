package main


import (
	"reddit_client_server/lozapi"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	client := lozapi.NewClient(lozapi.BaseUrl, &http.Client{
		Timeout: 10 * time.Second,
	})

	response, err := client.GetMosters()
	if err != nil {
		log.Fatal(err)
	}

	// for _, monster := range response.Data {
		fmt.Printf("%+v\n", response.Data)
	// }
}