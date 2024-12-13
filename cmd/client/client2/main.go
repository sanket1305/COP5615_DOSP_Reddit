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
// 3. createSubreddit "ufl"
// 4. JoinSubreddit "ufl"
// 5. getUsers
// 6. sendmessage
// 7. getsubrreddit
// 8. getfeed
// 9. makepost
// 10. getsubrreddit
// 11. getfeed
// 12. makecomment
// 13. getmessages

func main() {
	// use this fields to hold response received from apis
	// so that we can pass this values to subsequent api calls
	userName := ""
	subredditName := ""
	// recipient := ""

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
	}

	lozapi.Delay()

	// 3. create subreddit "ufl"
	subredditName = "ufl"
	response, err = client.CreateSubreddit(subredditName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", response.Message)

	lozapi.Delay()

	// 4. join subreddit "ufl"
	response, err = client.JoinSubreddit(userName, subredditName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", response.Message)

	lozapi.Delay()

	// 5. get list of users
	response_slice, err_slice = client.GetListOfAvailableUsers()
	if err_slice != nil {
		log.Fatal(err_slice)
	}

	if len(response_slice.Arr) == 0 {
		fmt.Println("No User available at the moment. You can create a new one though :)")
	} else {
		fmt.Println("Users available are:")
		for _, user := range response_slice.Arr {
			fmt.Printf("%+v\n", user)
		}
	}

	lozapi.Delay()

	// 6. sendmessage to each user
	if len(response_slice.Arr) == 0 {
		fmt.Println("No subreddits available at the moment. You can create a new one though :)")
	} else {
		for _, otherUser := range response_slice.Arr {
			if userName != otherUser {
				response, err = client.SendMessage(userName, otherUser, "Hello, Greetings!!!")
				if err != nil {
					log.Fatal(err)
				}

				fmt.Printf("%v\n", response.Message)
			}
		}
	}

	lozapi.Delay()

	// 7. getListofsubreddits (by now other clients must have added some more reddits and posts in them)
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
		
	}

	lozapi.Delay()

	// 8. getPosts (from any of the subreddits we have got above)
	// no response will be taken here, 
	// output will be printed inside the method GetFeed()
	err_list_post := client.GetFeed(subredditName)
	if err_list_post != nil {
		log.Fatal(err_list_post)
	}

	lozapi.Delay()

	// 9. Post in subreddit
	response, err = client.PostInSubreddit(userName, subredditName, "falana")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", response.Message)

	lozapi.Delay()

	// 10. getListofsubreddits (by now other clients must have added some more reddits and posts in them)
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
		
	}

	lozapi.Delay()

	// 11. getPosts (from any of the subreddits we have got above)
	// no response will be taken here, 
	// output will be printed inside the method GetFeed()
	err_list_post = client.GetFeed(subredditName)
	if err_list_post != nil {
		log.Fatal(err_list_post)
	}

	lozapi.Delay()

	// 12. add comment (for above post) -- post is hardcoded right now
	// we should now get random post above and make a comment on it
	response, err = client.CommentInSubreddit(userName, subredditName, "post1", "I agree with you")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", response.Message)

	lozapi.Delay()

	// 14. get messages
	response_inbox, err_inbox := client.CheckInbox(userName)
	if err_inbox != nil {
		log.Fatal(err_inbox)
	}

	for key, value := range response_inbox.Conversation {
		fmt.Printf("----- Displaying conversation with %s ------\n", key)
		for _, i := range value {
			if i[0] == key {
				fmt.Printf("Incoming... %s\n", i[1])
			} else {
				fmt.Printf("Ougoing... %s\n", i[1])
			}
		}
		fmt.Printf("----- End of Displaying conversation with %s ------\n", key)
	}

	lozapi.Delay()
}
