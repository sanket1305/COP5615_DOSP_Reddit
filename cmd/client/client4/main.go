package main

import (
	"fmt"
	"log"
	"net/http"
	"math/rand"
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
// 7. joinsubreddit
// 8. makepost
// 9. getsubredditlist
// 10. getfeed
// 11. makecomment
// 12. getuserlist
// 13. sendmessage

func main() {
	// use this fields to hold response received from apis
	// so that we can pass this values to subsequent api calls
	userName := ""
	subredditName := ""
	recipient := ""

	// create instance of client to call different APIs
	client := lozapi.NewClient(lozapi.BaseUrl, &http.Client{
		Timeout: 10 * time.Second,
	})

	// 1. register user
	response, err := client.RegisterUser()
	if err != nil {
		log.Fatal(err)
	}

	userName = response.Message
	fmt.Printf("Registration Successful!!! Your userName is %s\n", userName)

	lozapi.Delay()

	// 2. getListofsubreddits
	response_slice, err_slice := client.GetListOfAvailableSubreddits()
	if err_slice != nil {
		log.Fatal(err_slice)
	}

	if len(response_slice.Arr) == 0 {
		fmt.Println("No subreddits available at the moment. You can create a new one though :)")
	} else {
		fmt.Println("Available subreddits are :")
		for _, subre := range response_slice.Arr {
			fmt.Printf("%+v\n", subre)
		}
	// now we will try to pick some random subreddit
		// and interact with it in subsequent APIs
		// Create a new random generator with a seed based on the current time
		r := rand.New(rand.NewSource(time.Now().UnixNano()))

		// Generate a random index within the bounds of the slice
		randomIndex := r.Intn(len(response_slice.Arr))

		// This loop will make sure that the client will go to different subreddit
		for (subredditName == response_slice.Arr[randomIndex]) {
			randomIndex = r.Intn(len(response_slice.Arr))
		}

		// Select the random element from the slice
		subredditName = response_slice.Arr[randomIndex]
		fmt.Printf("Now we are looking at subreddit %s\n", subredditName)
	}

	lozapi.Delay()

	// 3. join subreddit "usa"
	response, err = client.JoinSubreddit(userName, subredditName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", response.Message)

	lozapi.Delay()

	// 4. getPosts (from any of the subreddits we have got above)
	// no response will be taken here, 
	// output will be printed inside the method GetFeed()
	err_list_post := client.GetFeed(subredditName)
	if err_list_post != nil {
		log.Fatal(err_list_post)
	}

	lozapi.Delay()

	// 5. leave subreddit
	response, err = client.LeaveSubreddit(userName, "usa")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", response.Message)

	// 6. create subreddit "football"
	subredditName = "football"
	response, err = client.CreateSubreddit(subredditName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", response.Message)
	fmt.Printf("Now we are looking at subreddit %s\n", subredditName)

	lozapi.Delay()

	// 7. join subreddit "football"
	response, err = client.JoinSubreddit(userName, subredditName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", response.Message)

	lozapi.Delay()

	// 8. Post in subreddit "football"
	response, err = client.PostInSubreddit(userName, subredditName, "Messi or Ronaldo?")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", response.Message)

	lozapi.Delay()

	// 9. getListofsubreddits
	response_slice, err_slice = client.GetListOfAvailableSubreddits()
	if err_slice != nil {
		log.Fatal(err_slice)
	}

	if len(response_slice.Arr) == 0 {
		fmt.Println("No subreddits available at the moment. You can create a new one though :)")
	} else {
		fmt.Println("Available subreddits are :")
		for _, subre := range response_slice.Arr {
			fmt.Printf("%+v\n", subre)
		}
		// now we will try to pick some random subreddit
		// and interact with it in subsequent APIs
		// Create a new random generator with a seed based on the current time
		r := rand.New(rand.NewSource(time.Now().UnixNano()))

		// Generate a random index within the bounds of the slice
		randomIndex := r.Intn(len(response_slice.Arr))

		// This loop will make sure that the client will go to different subreddit
		for (subredditName == response_slice.Arr[randomIndex]) {
			randomIndex = r.Intn(len(response_slice.Arr))
		}

		// Select the random element from the slice
		subredditName = response_slice.Arr[randomIndex]
		fmt.Printf("Now we are looking at subreddit %s\n", subredditName)
	}

	lozapi.Delay()

	// 10. getPosts (from any of the subreddits we have got above)
	// no response will be taken here, 
	// output will be printed inside the method GetFeed()
	err_list_post = client.GetFeed(subredditName)
	if err_list_post != nil {
		log.Fatal(err_list_post)
	}

	lozapi.Delay()

	// 11. add comment (for above post) -- post is hardcoded right now
	// we should now get random post above and make a comment on it
	fmt.Printf("Adding comment on post post1 in subreddit %s", subredditName)
	response, err = client.CommentInSubreddit(userName, subredditName, "post1", "I agree with you")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", response.Message)

	lozapi.Delay()

	// 12. get user list
	response_slice, err_slice = client.GetListOfAvailableUsers()
	if err_slice != nil {
		log.Fatal(err_slice)
	}

	if len(response_slice.Arr) == 0 {
		fmt.Println("No users available at the moment. You can be first one though :)")
	} else {
		for _, user := range response_slice.Arr {
			recipient = user
			fmt.Printf("%+v\n", user)
		}
	}

	lozapi.Delay()

	// 13. respond to message
	response, err = client.SendMessage(userName, recipient, "Heyy")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", response.Message)

	lozapi.Delay()
}