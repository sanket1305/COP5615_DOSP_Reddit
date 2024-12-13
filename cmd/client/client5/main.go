package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"reddit_client_server/lozapi"
)

// requests sequence

// 1. register
// 2. getListOfsubreddits
// 3. joinsubreddit
// 4. getfeed
// 5. leavsubreddit
// 6. createsubreddit
// 7. makepost
// 8. getsubredditlist
// 9. getfeed
// 10. makecomment
// 11. sendmessage

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
	username := "user1"

	// join subreddit
	response, err = client.JoinSubreddit(username, "USA")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", response.Message)

	// Post in subreddit
	response, err = client.PostInSubreddit(username, "USA", "falana falana")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", response.Message)

	// Comment in subreddit
	response, err = client.CommentInSubreddit(username, "USA", "post1", "No comments please!!!")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", response.Message)

	// get all posts in subreddit
	response, err = client.GetFeed("USA")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", response.Message)

	// Leave subreddit
	response, err = client.LeaveSubreddit(username, "USA")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", response.Message)
}
