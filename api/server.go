package main

import (
	"fmt"
	// "time"
	"encoding/json"
	"strconv"
	"log"
	"math/rand"
	"net/http"

    "github.com/gin-gonic/gin"
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

// for sending string responses
type Response struct {
    Message string `json:"message"`
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

type Upvote struct {
	User string
	Subreddit string
	Post string
	Comment string
}

type Downvote struct {
	User string
	Subreddit string
	Post string
	Comment string
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

type GetFeed struct {
	SubredditName string
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
	numComments int // total number of comments (direct comments + sub comments)
	Comments []string // list of direct comments
	AllComments map[string]*actor.PID // PIDs of all comments (direct comment + sub comments)
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
		state.numComments++
		newCommentID := fmt.Sprintf("com%d", state.numComments)
		commentProps := actor.PropsFromProducer(func() actor.Actor {return &CommentActor{CommentID: newCommentID, UserID: msg.User, Comment: msg.CommentTxt, SubComments: make([]*actor.PID, 0)}})
		commentPID := ctx.Spawn(commentProps)
		if state.AllComments == nil {
			state.AllComments = make(map[string]*actor.PID) // Initialize comments lazily
		}
		state.AllComments[newCommentID] = commentPID
		if msg.ParentComment == "" {
			state.Comments = append(state.Comments, newCommentID)
			fmt.Printf("%s commented on %s\n", msg.User, state.PostID)
		} else {
			pCommentID := state.AllComments[msg.ParentComment]
			ctx.Send(pCommentID, &MakeComment{User: msg.User, CommentTxt: msg.CommentTxt, CommentPID: commentPID, ParentComment: msg.ParentComment})
			fmt.Printf("%s subcommented on %s\n", msg.User, state.PostID)
		}
	
	case *DisplayComments:
		fmt.Println(state.numComments)
		fmt.Println(state.Comments)
	
	case *GetFeed:
		fmt.Println(state.UserID)
		fmt.Println(state.Content)
		fmt.Printf("The total number of comments + sub comments are %d\n", state.numComments)
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
			ctx.Send(postPID, &MakeComment{User: msg.User, CommentTxt: msg.CommentTxt, ParentComment: msg.ParentComment})
		} else {
			fmt.Printf("Post %s not found.\n", msg.Post)
		}

	case *DisplayComments:
		if postPID, ok := state.Posts[msg.PostName]; ok {
			ctx.Send(postPID, &DisplayComments{})
		} else {
			fmt.Printf("Post %s not found.\n", msg.PostName)
		}

	case *GetFeed:
		fmt.Printf("Fetching feed for subreddit %s\n", state.Name)
		for _, post := range state.Posts {
			ctx.Send(post, &GetFeed{})
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
			fmt.Printf("Subreddit %s not found. from EngineActor via JoinSubreddit\n", msg.SubredditName)
		}
	
	case *LeaveSubreddit:
		if subredditPID, ok := state.subreddits[msg.SubredditName]; ok {
			user := state.users[msg.UserID]
			ctx.Send(subredditPID, &LeaveSubreddit{UserID: msg.UserID})
			ctx.Send(user, &LeaveSubreddit{SubredditName: msg.SubredditName})
			// fmt.Printf("User %s left Subreddit %s. \n", msg.UserID, msg.SubredditName)
		} else {
			fmt.Printf("Subreddit %s not found. from EngineActor via LeaveSubreddit\n", msg.SubredditName)
		}
	
	case *PostinSubreddit:
		if subredditPID, ok := state.subreddits[msg.SubredditName]; ok {
			ctx.Send(subredditPID, &PostinSubreddit{Content: msg.Content, UserID: msg.UserID})
		} else {
			fmt.Printf("Subreddit %s not found. from EngineActor via PostinSubreddit\n", msg.SubredditName)
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
			ctx.Send(subredditPID, &MakeComment{User: msg.User, Post: msg.Post, CommentTxt: msg.CommentTxt, ParentComment: msg.ParentComment})
		} else {
			fmt.Printf("Subreddit %s not found. from EngineActor via MakeComment\n", msg.Subreddit)
		}
	
	case *Debug:
		// fmt.Println(state.msgs)
		fmt.Println(len(state.users))
		for index, value := range state.users {
			fmt.Printf("Index: %v, Value: %v\n", index, value)
		}
		fmt.Println(len(state.subreddits))
		for index, value := range state.subreddits {
			fmt.Printf("Index: %v, Value: %v\n", index, value)
		}
	
	case *DebugComments:
		subredditname := "sub2"
		postname := "post1"
		if subredditPID, ok := state.subreddits[subredditname]; ok {
			ctx.Send(subredditPID, &DisplayComments{PostName: postname})
		} else {
			fmt.Printf("Subreddit %s not found. from EngineActor via DebugComments\n", subredditname)
		}
	
	case *GetFeed:
		if subredditPID, ok := state.subreddits[msg.SubredditName]; ok {
			ctx.Send(subredditPID, &GetFeed{SubredditName: msg.SubredditName})
			fmt.Printf("Fetching feed for subreddit %s.\n", msg.SubredditName)
		} else {
			fmt.Printf("Subreddit %s not found. from EngineActor via GetFeed\n", msg.SubredditName)
		}
	
	case *Upvote:
		if subredditPID, ok := state.subreddits[msg.Subreddit]; ok {
			// if we receive non empty value in comment, then it's upvote for a comment
			// else it's upvote for a post 
			ctx.Send(subredditPID, &Upvote{Post: msg.Post, Comment: msg.Comment})
		} else {
			fmt.Printf("Subreddit %s not found.\n from EngineActor via Upvote", msg.Subreddit)
		}

	case *Downvote:
		if subredditPID, ok := state.subreddits[msg.Subreddit]; ok {
			ctx.Send(subredditPID, &Upvote{Post: msg.Post, Comment: msg.Comment})
		} else {
			fmt.Printf("Subreddit %s not found. from EngineActor via Downvote\n", msg.Subreddit)
		}

	default:
		log.Println("Unknown message type received")
	}
}

// ---------------------------- REST API FUNCTIONS START ---------------------------------

// Handler function to register new user
// func registerUser(c *gin.Context) {
//     c.JSON(http.StatusOK, users)
// }

// ---------------------------- REST API FUNCTIONS END ---------------------------------

func main() {
	userCount := 0
	// a central point for our actors and managing their lifecycle
	system := actor.NewActorSystem() 
	rootContext := system.Root

	// defining the properties of actor anf using it to spawn an instance
	engineProps := actor.PropsFromProducer(func() actor.Actor { return NewEngineActor() })
	enginePID := rootContext.Spawn(engineProps)

	// engine actor is ready for new msgs
	fmt.Println("Reddit server is live now!!!")

	// Create a Gin router
    router := gin.Default()

	// Define routes
    router.GET("/user/register", func(c *gin.Context) {
		userCount += 1
		username := "user" + strconv.Itoa(userCount)
		rootContext.Send(enginePID, &RegisterUser{Username: username})

		// Create an instance of the Response struct
		response := Response{
			Message: username + " is registered",
		}
	
		// Marshal the struct to JSON
		jsonData, err := json.Marshal(response)
		if err != nil {
			log.Fatalf("Error marshalling JSON: %v", err)
		}

		c.JSON(http.StatusOK, jsonData)
	})

	router.GET("/", func(c *gin.Context) {
		rootContext.Send(enginePID, &Debug{})

		response := Response{
			Message: "deubg called",
		}
	
		// Marshal the struct to JSON
		jsonData, err := json.Marshal(response)
		if err != nil {
			log.Fatalf("Error marshalling JSON: %v", err)
		}

		c.JSON(http.StatusOK, jsonData)
	})

	router.GET("/subreddit/create", func(c *gin.Context) {
		// req should contain {subredditname: string}
		subredditName := c.Query("subredditname")

		rootContext.Send(enginePID, &CreateSubreddit{Name: subredditName})

		response := Response{
			Message: "subreddit created",
		}
	
		// Marshal the struct to JSON
		jsonData, err := json.Marshal(response)
		if err != nil {
			log.Fatalf("Error marshalling JSON: %v", err)
		}

		c.JSON(http.StatusOK, jsonData)
	})

	router.GET("/subreddit/join", func(c *gin.Context) {
		// req should contain {subredditname: string}
		subredditName := c.Query("subredditname")
		userName := c.Query("username")

		// fmt.Println(subredditName)
		// fmt.Println(userName)

		rootContext.Send(enginePID, &JoinSubreddit{UserID: userName, SubredditName: subredditName})	

		response := Response{
			Message: "subreddit joined",
		}
	
		// Marshal the struct to JSON
		jsonData, err := json.Marshal(response)
		if err != nil {
			log.Fatalf("Error marshalling JSON: %v", err)
		}

		c.JSON(http.StatusOK, jsonData)
	})

	router.GET("/subreddit/leave", func(c *gin.Context) {
		// req should contain {subredditname: string}
		subredditName := c.Query("subredditname")
		userName := c.Query("username")

		fmt.Println(subredditName)
		fmt.Println(userName)

		rootContext.Send(enginePID, &LeaveSubreddit{UserID: userName, SubredditName: subredditName})

		response := Response{
			Message: "subreddit left",
		}
	
		// Marshal the struct to JSON
		jsonData, err := json.Marshal(response)
		if err != nil {
			log.Fatalf("Error marshalling JSON: %v", err)
		}

		c.JSON(http.StatusOK, jsonData)
	})

    // router.POST("/user/karma", getUserKarma)
	// router.POST("/user/messaging", sendMessage)
	// // router.GET("/user/listsubeddits", addUser)
	// Done... router.POST("/subreddit/create", createSubred)
	// Done... router.POST("/subreddit/join", joinSubred)
	// Done... router.POST("/subreddit/leave", leaveSubred)
	// router.POST("/subreddit/post", postSubred)
	// router.POST("/subreddit/comment", commentSubred)
	// router.POST("/subreddit/feed", feedSubred)
	// // router.GET("/subreddit/listusers", addUser)
	// router.POST("/post/upvote", upvotePost)
	// router.POST("/post/downvote", downvotePost)

    // Start the server on port 8080
    router.Run("localhost:8080")
}