// to do

// Register Account (Done)
// Create subreddit (Done)
// Join subreddit (Done) added check for duplicate entries during simulation
// Leave subreddit (Done) added check to verify if the pair exists
// Post in subreddit (Done)
// Comment in subreddit
// 	Hierarchial view
// Upvote, Downvote, compute Karma (Done)
// get feed of posts
// get lists of direct messages; reply to direct messages (Done)
	// engine actor will maintain one feild, which will store all the messages in map
	// the map key will hold [user1, user2]... it will optimize the performance, to retriev the chatsin O(1)
	// the value will hold slice[slice].... where each element will indicate, what's the message and it's send by which user

// implement simulator
// 	simulate as many users as you can (10 at start)
// 	simulate periods of live connection and disconnection for users
// 	simulate a zipf ditribution on the number of sub-reddit members.
// 		for account with a lot of subscribers, increase the number of posts.
// 		Make some of these messages re-posts

// compute karma
// for each upvote/downvote receoived +/-1
// for each new post +5 karma


package main

import (
	"fmt"
	"time"
	"log"
	"math/rand"

	"github.com/asynkron/protoactor-go/actor"
)

// function to check if value exists in slice
func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// for debuging
type Debug struct {}

type DebugComments struct {}

// Message types
type RegisterUser struct {
	Username string
}

type CreateSubreddit struct {
	Name string
}

type JoinSubreddit struct {
	UserID       	string
	SubredditName 	string
}

type LeaveSubreddit struct {
	UserID       	string
	SubredditName	string
}

type PostinSubreddit struct{
	PostID			string
	Content			string
	UserID			string
	SubredditName	string
	upvotes			int
	downvotes		int
}

type UpdateKarma struct {
    UserID    	string
    Karma		int
}

type DirectMessage struct {
	User1	string
	User2	string
	Message	string
}

type MakeComment struct {
	CommentPID	*actor.PID
	User		string
	Subreddit	string
	Post		string
	CommentTxt	string	
	ParentComment	string	
}

type DisplayComments struct {
	PostName string
}

// user actor
type UserActor struct {
	ID       		string
	Username 		string
	Karma    		int
	SubredditList	[]string
}

// behavior for user actor, which will act as listener
func (state *UserActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *RegisterUser:
		state.Username = msg.Username
		state.Karma = 0
		fmt.Printf("%s registered.\n", state.Username)

	case *JoinSubreddit:
		if !contains(state.SubredditList, msg.SubredditName) {
			state.SubredditList = append(state.SubredditList, msg.SubredditName)
			fmt.Printf("%s joined subreddit %s.\n", state.ID, msg.SubredditName)
			// fmt.Println(state.SubredditList)
		}
	
	case *LeaveSubreddit:
		currsubreddits := state.SubredditList
		// fmt.Println("Before") // DEBUG
		// fmt.Println(currsubreddits) // DEBUG
		for i, v := range currsubreddits {
			
			if v == msg.SubredditName {
				// Remove the element at index i
				currsubreddits = append(currsubreddits[:i], currsubreddits[i+1:]...)
				break
			}
		}
		state.SubredditList = currsubreddits
		fmt.Printf("%s left subreddit %s. \n", state.ID, msg.SubredditName)
		// fmt.Println("After") // DEBUG
		// fmt.Println(currsubreddits) // DEBUG
	
	case *UpdateKarma:
        if msg.UserID == state.ID { // Ensure it's for this user
            state.Karma += msg.Karma
            fmt.Printf("%s's Karma has been updated to %d.\n", state.ID, state.Karma)
        }
	}
}

type CommentActor struct {
	CommentID string
	UserID string
	Comment string
	SubComments []*actor.PID
}

func (state *CommentActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *MakeComment:
		state.UserID = msg.User
		state.Comment = msg.CommentTxt
		state.SubComments = append(state.SubComments, msg.CommentPID)
	}
}

type PostActor struct {
	PostID string
	UserID string
	Content string
	numComments int
	Comments []string
	AllComments map[string]*actor.PID
	upvotes int
	downvotes int
}

func (state *PostActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *PostinSubreddit:
		state.PostID = msg.PostID
		state.Content = msg.Content
		state.UserID = msg.UserID
		state.numComments = 0
		// state.AllComments = make(map[string]*actor.PID)
		state.Comments = make([]string, 0)
		fmt.Printf("Post %s created.\n", state.PostID)

	case *MakeComment:
		newCommentID := fmt.Sprintf("com%d", state.numComments + 1)
		commentProps := actor.PropsFromProducer(func() actor.Actor {return &CommentActor{CommentID: newCommentID, UserID: msg.User, Comment: msg.CommentTxt, SubComments: make([]*actor.PID, 0)}})
		commentPID := ctx.Spawn(commentProps)
		if state.AllComments == nil {
			state.AllComments = make(map[string]*actor.PID) // Initialize comments lazily
		}
		state.AllComments[newCommentID] = commentPID
		if msg.ParentComment == "" {
			state.Comments = append(state.Comments, newCommentID)
		} else {
			pCommentID := state.AllComments[msg.ParentComment]
			ctx.Send(pCommentID, &MakeComment{User: msg.User, CommentTxt: msg.CommentTxt, CommentPID: commentPID})
		}
	
	case *DisplayComments:
		fmt.Println(state.numComments)
		fmt.Println(state.Comments)
	}
}

// Subreddit Actor
type SubredditActor struct {
	Name  string
	numPosts int
	Posts map[string]*actor.PID // List of posts in the subreddit.
	UserList []string
	engine *actor.PID
}

// behavior for sub reddit, to listen to incoming messages
func (state *SubredditActor) Receive(ctx actor.Context) {   // , rootContext *actor.RootContext, enginePID *actor.PID
	switch msg := ctx.Message().(type) {
	case *CreateSubreddit:
		state.Name = msg.Name
		state.numPosts = 0
		// state.Posts = make(map[string]*actor.PID) // Initialize Posts map here
		fmt.Printf("Subreddit %s created.\n", state.Name)

	case *JoinSubreddit:
		if !contains(state.UserList, msg.UserID) {
			state.UserList = append(state.UserList, msg.UserID)
		}
	
	case *LeaveSubreddit:
		currusers := state.UserList
		// fmt.Println("Before") // DEBUG
		// fmt.Println(currusers) // DEBUG
		for i, v := range currusers {
			
			if v == msg.UserID {
				// Remove the element at index i
				currusers = append(currusers[:i], currusers[i+1:]...)
				break
			}
		}
		state.UserList = currusers
		// fmt.Println("After")  // DEBUG
		// fmt.Println(currusers)  // DEBUG
	
	case *PostinSubreddit:
		if state.Posts == nil {
			state.Posts = make(map[string]*actor.PID) // Initialize Posts lazily
		}
		numUsers := len(state.UserList)
		
		for i, v := range state.UserList {

			if v == msg.UserID {
				state.numPosts++
				nPosts := fmt.Sprintf("post%d", state.numPosts)
				upv := rand.Intn(numUsers)
				dnv := rand.Intn(numUsers - upv)
				postProps := actor.PropsFromProducer(func() actor.Actor { return &PostActor{PostID: nPosts ,Content: msg.Content, UserID: msg.UserID, upvotes: upv, downvotes: dnv} })
				postid := ctx.Spawn(postProps) 
				state.Posts[nPosts] = postid
				fmt.Printf("%s posted in subreddit %s. \n", state.UserList[i], state.Name)

                ctx.Send(state.engine, &UpdateKarma{UserID: msg.UserID, Karma: upv-dnv})
				break
			}
		}
	
	case *MakeComment:
		if postPID, ok := state.Posts[msg.Post]; ok {
			ctx.Send(postPID, &MakeComment{User: msg.User, CommentTxt: msg.CommentTxt})
		} else {
			fmt.Printf("Post %s not found.\n", msg.Post)
		}

	case *DisplayComments:
		if postPID, ok := state.Posts[msg.PostName]; ok {
			ctx.Send(postPID, &DisplayComments{})
		} else {
			fmt.Printf("Post %s not found.\n", msg.PostName)
		}
	}

}

// Engine Actor (Orchestrator)
type EngineActor struct {
	users      	map[string]*actor.PID
	subreddits 	map[string]*actor.PID
	msgs 	   	map[[2]string][][]string
}

// function to return pointer to engine actor with empty map of users and subreddits
func NewEngineActor() *EngineActor {
	return &EngineActor{
		users:      make(map[string]*actor.PID),
		subreddits: make(map[string]*actor.PID),
		msgs:		make(map[[2]string][][]string),
	}
}

// behavior for engineActor, to listen for incoming messages
func (state *EngineActor) Receive(ctx actor.Context) {
	// checking msg type, based on msg struct
	switch msg := ctx.Message().(type) {

	case *RegisterUser:
		userProps := actor.PropsFromProducer(func() actor.Actor { return &UserActor{ID: msg.Username} })
		userPID := ctx.Spawn(userProps)
		state.users[msg.Username] = userPID
		// fmt.Printf("User %s registered.\n", msg.Username)
	
	case *CreateSubreddit:
		subredditProps := actor.PropsFromProducer(func() actor.Actor { return &SubredditActor{Name: msg.Name, engine: ctx.Self()} })
		subredditPID := ctx.Spawn(subredditProps)
		state.subreddits[msg.Name] = subredditPID
		// fmt.Printf("Subreddit %s created.\n", msg.Name)
	
	case *JoinSubreddit:
		if subredditPID, ok := state.subreddits[msg.SubredditName]; ok {
			user := state.users[msg.UserID]
			ctx.Send(subredditPID, &JoinSubreddit{UserID: msg.UserID})
			ctx.Send(user, &JoinSubreddit{SubredditName: msg.SubredditName})
			// fmt.Printf("User %s joined subreddit %s.\n", msg.UserID, msg.SubredditName)
		} else {
			fmt.Printf("Subreddit %s not found.\n", msg.SubredditName)
		}
	
	case *LeaveSubreddit:
		if subredditPID, ok := state.subreddits[msg.SubredditName]; ok {
			user := state.users[msg.UserID]
			ctx.Send(subredditPID, &LeaveSubreddit{UserID: msg.UserID})
			ctx.Send(user, &LeaveSubreddit{SubredditName: msg.SubredditName})
			// fmt.Printf("User %s left Subreddit %s. \n", msg.UserID, msg.SubredditName)
		} else {
			fmt.Printf("Subreddit %s not found.\n", msg.SubredditName)
		}
	
	case *PostinSubreddit:
		if subredditPID, ok := state.subreddits[msg.SubredditName]; ok {
			ctx.Send(subredditPID, &PostinSubreddit{Content: msg.Content, UserID: msg.UserID})
		} else {
			fmt.Printf("Subreddit %s not found.\n", msg.SubredditName)
		}
	
	case *UpdateKarma:
		user := state.users[msg.UserID]
		// fmt.Printf("%s's karma is received in Engine\n", msg.UserID)
		ctx.Send(user, &UpdateKarma{UserID:msg.UserID, Karma: msg.Karma})

	case *DirectMessage:
		// create keys for searching in
		key := [2]string{msg.User1, msg.User2}
		key_rev := [2]string{msg.User2, msg.User1}

		if value, ok := state.msgs[key]; ok {
			newMessage := []string{msg.User1, msg.Message}
			value = append(value, newMessage)
			state.msgs[key] = value

		} else if value, ok := state.msgs[key_rev]; ok {
			newMessage := []string{msg.User1, msg.Message}
			value = append(value, newMessage)
			state.msgs[key] = value
		} else {
			newMessage := []string{msg.User1, msg.Message}
			value := make([][]string, 0)
			value = append(value, newMessage)
			state.msgs[key] = value
		}
	
	case *MakeComment:
		if subredditPID, ok := state.subreddits[msg.Subreddit]; ok {
			ctx.Send(subredditPID, &MakeComment{User: msg.User, Post: msg.Post, CommentTxt: msg.CommentTxt})
		} else {
			fmt.Printf("Subreddit %s not found.\n", msg.Subreddit)
		}
	
	case *Debug:
		fmt.Println(state.msgs)
	
	case *DebugComments:
		subredditname := "sub1"
		postname := "post1"
		if subredditPID, ok := state.subreddits[subredditname]; ok {
			ctx.Send(subredditPID, &DisplayComments{PostName: postname})
		} else {
			fmt.Printf("Subreddit %s not found.\n", subredditname)
		}

	default:
		log.Println("Unknown message type received")
	}
}		

// // function to simulate our reddit engine for multi-user scenario
// func simulateUsers(rootContext *actor.RootContext, enginePID *actor.PID, numUsers int) {

// 	//Approach:
// 	// create 10 users (Done)
// 	// create 5 subreddits (Done)
// 	// for each user, use random function (5 times) and assign the user to random subreedits among available ones (Done)
// 	// run for loop for 20 times, make random posts by users in subreddits
// 	// run for loop for 20 times, make random comments (non-hierarchial0)
// 	// upvote 100 times, random posts
// 	// downvote 100 times random posts
// 	// make 2 users leave the subreddit
// 	// get feed for all 5 subreddits 

// 	// creating 10 users
// 	for i := 0; i < numUsers; i++ {
//         username := fmt.Sprintf("user%d", i+1)

//         // Register the user.
//         rootContext.Send(enginePID, &RegisterUser{Username: username})
// 	}

// 	// creating 5 subredits
// 	for i := 0; i < 5; i++ {
// 		subredditName := fmt.Sprintf("sub%d", i+1)
// 		rootContext.Send(enginePID, &CreateSubreddit{Name: subredditName})
// 	}

// 	// join subreddits
// 	for i := 0; i < 10; i++ {
// 		username := fmt.Sprintf("user%d", i+1)
// 		for j := 0; j < 5; j++ {
// 			subredditName := fmt.Sprintf("sub%d", rand.Intn(5)+1) // Randomly choose from 5 subreddits.

// 			rootContext.Send(enginePID, &JoinSubreddit{UserID: username, SubredditName: subredditName})
// 		}
// 	}

// 	// users posts in a subreddit
// 	// for i := 0; i < 10; i++ {
// 	// 	username := fmt.Sprintf("user%d", rand.Intn(10)+1)
// 	// 	subredditname := fmt.Sprintf("sub%d", rand.Intn(5)+1) // Randomly choose from user's subreddits.
// 	// 	content := fmt.Sprintf("user %s posted this content in subreddit %s", username, subredditname)
// 	// 	rootContext.Send(enginePID, &PostinSubreddit{UserID: username, SubredditName: subredditname, Content: content})
// 	// }

// 	username := "user1"
// 	subredditname := "sub1"
// 	content := "user1 posted this content in subreddit1"
// 	rootContext.Send(enginePID, &PostinSubreddit{UserID: username, SubredditName: subredditname, Content: content})

// 	// users leaves randomly
// 	for i:= 0; i< 5; i++ {
// 		username := fmt.Sprintf("user%d", rand.Intn(10)+1)  
// 		subredditname:= fmt.Sprintf("sub%d", rand.Intn(5)+1)   // Randomly choose from user's subreddits.

// 		// time.Sleep(3 * time.Second)
// 		// fmt.Println(username)
// 		// fmt.Println(subredditname)
// 		rootContext.Send(enginePID, &LeaveSubreddit{UserID: username, SubredditName: subredditname})
// 	}

// 	// send messages
// 	username1 := "user1"
// 	username2 := "user2"
// 	content = "user1 sending hi to user2"

// 	rootContext.Send(enginePID, &DirectMessage{User1: username1, User2: username2, Message: content})
// 	// time.Sleep(5 * time.Second)

// 	rootContext.Send(enginePID, &Debug{})

// 	// comment on posts
// 	username = "user1"
// 	subreddit := "sub1"
// 	post := "post1"
// 	// comment := "com1"
// 	commentTxt := "This is 1st comment."

// 	rootContext.Send(enginePID, &MakeComment{User: username, Subreddit: subreddit, Post: post, CommentTxt: commentTxt})

// 	// parentComment := "com1"
// 	// commentTxt = "This is 2nd comment."

// 	// rootContext.Send(enginePID, &MakeComment{User: username, Subreddit: subreddit, Post: post, CommentTxt: commentTxt, ParentComment: parentComment})

// 	// parentComment = "com1"
// 	// commentTxt = "This is 3rd comment."

// 	// rootContext.Send(enginePID, &MakeComment{User: username, Subreddit: subreddit, Post: post, CommentTxt: commentTxt, ParentComment: parentComment})

// 	// parentComment = "com2"
// 	// commentTxt = "This is 4th comment."

// 	// rootContext.Send(enginePID, &MakeComment{User: username, Subreddit: subreddit, Post: post, CommentTxt: commentTxt, ParentComment: parentComment})

// 	rootContext.Send(enginePID, &DebugComments{})

// }




func simulateUsers(rootContext *actor.RootContext, enginePID *actor.PID, numUsers int) {

	// creating 10 users
	for i := 1; i <= numUsers; i++ {
		username := fmt.Sprintf("user%d", i)

		// Register the user.
		rootContext.Send(enginePID, &RegisterUser{Username: username})
	}

	// creating 5 subredits
	for i := 1; i <= 5; i++ {
		subredditName := fmt.Sprintf("sub%d", i)
		rootContext.Send(enginePID, &CreateSubreddit{Name: subredditName})
	}

	// join subreddit
	for i := 1; i <= 10; i++ {
		username := fmt.Sprintf("user%d", i)
		for j := 1; j <= i && j<=5; j++ {
			subredditName := fmt.Sprintf("sub%d", j) // Randomly choose from 5 subreddits.
			rootContext.Send(enginePID, &JoinSubreddit{UserID: username, SubredditName: subredditName})
		}
	}

	// post twice in each subreddit respectively
	for i := 1; i < 10; i++ {
		username := fmt.Sprintf("user%d", i)
		for j := 1; j <= i && j<=5; j++ {
			subredditname := fmt.Sprintf("sub%d", j) // choose from user's subreddits.
			content1 := fmt.Sprintf("%s posted this content in subreddit %s", username, subredditname)
			content2 := fmt.Sprintf("again %s posted this content in subreddit %s", username, subredditname)
			rootContext.Send(enginePID, &PostinSubreddit{UserID: username, SubredditName: subredditname, Content: content1})
			rootContext.Send(enginePID, &PostinSubreddit{UserID: username, SubredditName: subredditname, Content: content2})
		}
	}

	// users leaves randomly
	u1 := rand.Intn(9)
	u2 := u1+1
	username1 := fmt.Sprintf("user%d", u1)
	username2 := fmt.Sprintf("user%d", u2)
	subredditname1 := fmt.Sprintf("sub%d", rand.Intn(min(u1, 5))) // Randomly choose from user's subreddits.
	subredditname2 := fmt.Sprintf("sub%d", rand.Intn(min(u2, 5))) // Randomly choose from user's subreddits.
	// time.Sleep(3 * time.Second)
	// fmt.Println(username)
	// fmt.Println(subredditname)
	rootContext.Send(enginePID, &LeaveSubreddit{UserID: username1, SubredditName: subredditname1})
	rootContext.Send(enginePID, &LeaveSubreddit{UserID: username2, SubredditName: subredditname2})

	// send messages
	// username1 := "user1"
	// username2 := "user2"
	content := "user1 sending hi to user2"

	rootContext.Send(enginePID, &DirectMessage{User1: username1, User2: username2, Message: content})
	// time.Sleep(5 * time.Second)

	rootContext.Send(enginePID, &Debug{})

	time.Sleep(5 * time.Second)
	// comment on posts
	username := "user3"
	subreddit := "sub2"
	post := "post1"
	// comment := "com1"
	commentTxt := "This is 1st comment."

	rootContext.Send(enginePID, &MakeComment{User: username, Subreddit: subreddit, Post: post, CommentTxt: commentTxt})
	
	time.Sleep(2 * time.Second)
	rootContext.Send(enginePID, &DebugComments{})
	
	// for i := 0; i<=10; i++


}




func main() {

	// a central point for our actors and managing their lifecycle
	system := actor.NewActorSystem() 
	rootContext := system.Root

	// defining the properties of actor anf using it to spawn an instance
	engineProps := actor.PropsFromProducer(func() actor.Actor { return NewEngineActor() })
	enginePID := rootContext.Spawn(engineProps)

    // initiate simulator
    simulateUsers(rootContext, enginePID, 10)

	// delay main thread for some time,
	// allow other processes to finish
	time.Sleep(8 * time.Second) 
}