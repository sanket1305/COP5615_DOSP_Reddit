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
// 3. createSubreddit "USA"
// 4. JoinSubreddit
// 5. MakePost
// 6. getListOfsubreddits
// 7. getPosts from other subreddit
// 8. Add comment
// 9. get post from subreddit
// 10. get messages
// 11. respond to message
// 12. get feed from current subreddit
// 13. get member list from other subreddit
// 14. get posts from other subreddit
// 15. get messages

func main() {
	// create instance of client to call different APIs
	client := lozapi.NewClient(lozapi.BaseUrl, &http.Client{
		Timeout: 10 * time.Second,
	})

	// 1. register user
	response, err := client.RegisterUser()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", response.Message)

	// 2. getListofsubreddits
	response_slice, err_slice := client.GetListOfAvailableSubreddits()
	if err_slice != nil {
		log.Fatal(err_slice)
	}

	if len(response_slice.Arr) == 0 {
		fmt.Println("No subreddits available at the moment. You can create a new one though :)")
	} else {
		for _, subre := range response_slice.Arr {
			fmt.Printf("%+v\n", subre)
		}
	}

	// 3. create subreddit "USA"
	response, err = client.CreateSubreddit("USA")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", response.Message)
	username := "user1"

	// 4. join subreddit "USA"
	response, err = client.JoinSubreddit(username, "USA")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", response.Message)

	// 5. Post in subreddit "USA"
	response, err = client.PostInSubreddit(username, "USA", "falana")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", response.Message)

	// 5. Post in subreddit "USA"
	response, err = client.PostInSubreddit(username, "USA", "chin")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", response.Message)

	// 6. getListofsubreddits (by now other clients must have added some more reddits and posts in them)
	// this code is same code in 2.
	response_slice, err_slice = client.GetListOfAvailableSubreddits()
	if err_slice != nil {
		log.Fatal(err_slice)
	}

	if len(response_slice.Arr) == 0 {
		fmt.Println("No subreddits available at the moment. You can create a new one though :)")
	} else {
		for _, subre := range response_slice.Arr {
			fmt.Printf("%+v\n", subre)
		}
	}

	// 7. getPosts (from any of the subreddits we have got above)

	// 8. add comment (for above post)

	// 9. repeat 6 and 7

	// 10. get message (by now this should have received some message from client2 as it sends msg to all availble clients)

	// 11. respond to message
	
	// 12. repeat 6

	// 13. get member list from other subreddit

	// 14. repeat 6 and 7

	// 15. get messages

	// Comment in subreddit
	response, err = client.CommentInSubreddit(username, "USA", "post1", "No comments please!!!")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", response.Message)

	// get all posts in subreddit
	err_list_post := client.GetFeed("USA")
	if err_list_post != nil {
		log.Fatal(err_list_post)
	}

	// fmt.Printf("%v\n", response_list_post)

	// Leave subreddit
	response, err = client.LeaveSubreddit(username, "USA")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", response.Message)
}
