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
// 3. createSubreddit "USA"
// 4. JoinSubreddit "USA"
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

func delay() {
	time.Sleep(2 * time.Second)
}

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

	delay()

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

	delay()

	// 3. create subreddit "USA"
	subredditName = "USA"
	response, err = client.CreateSubreddit(subredditName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", response.Message)

	delay()

	// 4. join subreddit "USA"
	response, err = client.JoinSubreddit(userName, subredditName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", response.Message)

	delay()

	// 5. Post in subreddit "USA"
	response, err = client.PostInSubreddit(userName, subredditName, "falana")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", response.Message)

	delay()

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
		// now we will try to pick some random subreddit
		// and interact with it in subsequent APIs
		// Create a new random generator with a seed based on the current time
		r := rand.New(rand.NewSource(time.Now().UnixNano()))

		// Generate a random index within the bounds of the slice
		randomIndex := r.Intn(len(response_slice.Arr))

		// Select the random element from the slice
		subredditName = response_slice.Arr[randomIndex]
	}

	delay()

	// 7. getPosts (from any of the subreddits we have got above)
	// no response will be taken here, 
	// output will be printed inside the method GetFeed()
	err_list_post := client.GetFeed(subredditName)
	if err_list_post != nil {
		log.Fatal(err_list_post)
	}

	delay()

	// 8. add comment (for above post) -- post is hardcoded right now
	// we should now get random post above and make a comment on it
	response, err = client.CommentInSubreddit(userName, subredditName, "post1", "I agree with you")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", response.Message)

	delay()

	// 9. repeat 6 and 7
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
		// now we will try to pick some random subreddit
		// and interact with it in subsequent APIs
		// Create a new random generator with a seed based on the current time
		r := rand.New(rand.NewSource(time.Now().UnixNano()))

		// Generate a random index within the bounds of the slice
		randomIndex := r.Intn(len(response_slice.Arr))

		// Select the random element from the slice
		subredditName = response_slice.Arr[randomIndex]
	}

	delay()

	// 7. getPosts (from any of the subreddits we have got above)
	// no response will be taken here, 
	// output will be printed inside the method GetFeed()
	err_list_post = client.GetFeed(subredditName)
	if err_list_post != nil {
		log.Fatal(err_list_post)
	}

	delay()

	// 10. get message (by now this should have received some message from client2 as it sends msg to all availble clients)
	response_inbox, err_inbox := client.CheckInbox(userName)
	if err_inbox != nil {
		log.Fatal(err_inbox)
	}

	for key, value := range response_inbox.Conversation {
		recipient = value[0][0]
		fmt.Printf("----- Displaying conversation with %s ------", key)
		for _, i := range value {
			if i[0] == key {
				fmt.Printf("Incoming... %s", i[1])
			} else {
				fmt.Printf("Ougoing... %s", i[2])
			}
		}
	}

	delay()

	// 11. respond to message
	response, err = client.SendMessage(userName, recipient, "Heyy")
	if err != nil {
		log.Fatal(err_inbox)
	}

	fmt.Printf("%v\n", response.Message)

	response_inbox, err_inbox = client.CheckInbox("user1")
	if err_inbox != nil {
		log.Fatal(err_inbox)
	}

	delay()
	
	// 12. repeat 6
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

	delay()

	// 13. get user list
	response_slice, err_slice = client.GetListOfAvailableUsers()
	if err_slice != nil {
		log.Fatal(err_slice)
	}

	if len(response_slice.Arr) == 0 {
		fmt.Println("No users available at the moment. You can be first one though :)")
	} else {
		for _, subre := range response_slice.Arr {
			fmt.Printf("%+v\n", subre)
		}
	}

	delay()

	// 14. get messages

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

	delay()

	// 15. leave subreddit
	response, err = client.LeaveSubreddit(userName, "USA")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", response.Message)
}
