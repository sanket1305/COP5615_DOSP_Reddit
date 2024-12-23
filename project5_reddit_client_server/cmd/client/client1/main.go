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
// 3. createSubreddit "usa"
// 4. JoinSubreddit "usa"
// 5. MakePost "USA election results are out"
// 6. getListOfsubreddits
// 10. get feed
// 11. get messages
// 12. respond to message "Heyyy"
// 13. getlistofsubreddits and choose different one
// 14. get member list
// 15. get messages
// 16. leave subreddit

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
	}

	lozapi.Delay()

	// 3. create subreddit "usa"
	subredditName = "usa"
	response, err = client.CreateSubreddit(subredditName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", response.Message)

	lozapi.Delay()

	// 4. join subreddit "usa"
	response, err = client.JoinSubreddit(userName, subredditName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", response.Message)

	lozapi.Delay()

	// 5. Post in subreddit "usa"
	fmt.Printf("Making post in subreddit %s\n", subredditName)
	response, err = client.PostInSubreddit(userName, subredditName, "USA election results are out")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", response.Message)

	lozapi.Delay()

	// 6. getListofsubreddits (by now other clients must have added some more subreddits and posts in them)
	// this code is same code in 2.
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

	// // 7. getPosts (from any of the subreddits we have got above)
	// // no response will be taken here, 
	// // output will be printed inside the method GetFeed()
	// err_list_post := client.GetFeed(subredditName)
	// if err_list_post != nil {
	// 	log.Fatal(err_list_post)
	// }

	// lozapi.Delay()

	// // 8. add comment (for above post) -- post is hardcoded right now
	// // we should now get random post above and make a comment on it
	// fmt.Printf("Adding comment on post post1 in subreddit %s", subredditName)
	// response, err = client.CommentInSubreddit(userName, subredditName, "post1", "I agree with you")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%s\n", response.Message)

	// lozapi.Delay()

	// // 9. getListofsubreddits (by now other clients must have added some more reddits and posts in them)
	// // this code is same code in 2.
	// response_slice, err_slice = client.GetListOfAvailableSubreddits()
	// if err_slice != nil {
	// 	log.Fatal(err_slice)
	// }

	// if len(response_slice.Arr) == 0 {
	// 	fmt.Println("No subreddits available at the moment. You can create a new one though :)")
	// } else {
	// 	fmt.Println("Available subreddits are :")
	// 	for _, subre := range response_slice.Arr {
	// 		fmt.Printf("%+v\n", subre)
	// 	}
	// 	// now we will try to pick some random subreddit
	// 	// and interact with it in subsequent APIs
	// 	// Create a new random generator with a seed based on the current time
	// 	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 	// Generate a random index within the bounds of the slice
	// 	randomIndex := r.Intn(len(response_slice.Arr))

	// 	// This loop will make sure that the client will go to different subreddit
	// 	for (subredditName == response_slice.Arr[randomIndex]) {
	// 		randomIndex = r.Intn(len(response_slice.Arr))
	// 	}

	// 	// Select the random element from the slice
	// 	subredditName = response_slice.Arr[randomIndex]
	// 	fmt.Printf("Now we are looking at subreddit %s\n", subredditName)
	// }

	// lozapi.Delay()

	// 10. getPosts (from any of the subreddits we have got above)
	// no response will be taken here, 
	// output will be printed inside the method GetFeed()
	err_list_post := client.GetFeed(subredditName)
	if err_list_post != nil {
		log.Fatal(err_list_post)
	}

	lozapi.Delay()

	// 11. get message (by now this should have received some message from client2 as it sends msg to all availble clients)
	response_inbox, err_inbox := client.CheckInbox(userName)
	if err_inbox != nil {
		log.Fatal(err_inbox)
	}

	for key, value := range response_inbox.Conversation {
		recipient = key
		fmt.Printf("----- Displaying conversation with %s ------\n", key)
		for _, i := range value {
			if i[0] == key {
				// recipient = i[1]
				fmt.Printf("Incoming... %s\n", i[1])
			} else {
				// recipient = i[0]
				fmt.Printf("Ougoing... %s\n", i[2])
			}
		}
		fmt.Printf("----- End of Displaying conversation with %s ------\n", key)
	}

	lozapi.Delay()

	// 12. respond to message
	response, err = client.SendMessage(userName, recipient, "Heyy")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", response.Message)

	lozapi.Delay()
	
	// 13. getListofsubreddits (by now other clients must have added some more reddits and posts in them)
	// this code is same code in 2.
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

	// 14. get user list
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

	lozapi.Delay()

	// 15. get messages

	response_inbox, err_inbox = client.CheckInbox(userName)
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

	// 16. leave subreddit
	response, err = client.LeaveSubreddit(userName, "usa")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", response.Message)
}
