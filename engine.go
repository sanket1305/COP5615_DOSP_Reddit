// to do

// Register Account (Done)
// Create subreddit
// Join subreddit
// Leave subreddit
// Post in subreddit
// Comment in subreddit
// 	Hierarchial view
// Upvote, Downvote, compute Karma
// get feed of posts
// get lists of direct messages; reply to direct messages

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

	"github.com/asynkron/protoactor-go/actor"
)

// Message types
type RegisterUser struct {
	Username string
}

// user actor
type UserActor struct {
	ID       string
	Username string
	Karma    int
}

// behavior for user actor, which will act as listener
func (state *UserActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *RegisterUser:
		state.Username = msg.Username
		fmt.Printf("User %s registered.\n", state.Username)
	}
}

// Engine Actor (Orchestrator)
type EngineActor struct {
	users      map[string]*actor.PID
	subreddits map[string]*actor.PID
}

// function to return pointer to engine actor with empty map of users and subreddits
func NewEngineActor() *EngineActor {
	return &EngineActor{
		users:      make(map[string]*actor.PID),
		subreddits: make(map[string]*actor.PID),
	}
}

// behaviour for engineActor, to listen for incoming messages
func (state *EngineActor) Receive(ctx actor.Context) {
	// checking msg type, based on msg struct
	switch msg := ctx.Message().(type) {

	case *RegisterUser:
		userProps := actor.PropsFromProducer(func() actor.Actor { return &UserActor{ID: msg.Username} })
		userPID := ctx.Spawn(userProps)
		state.users[msg.Username] = userPID
		// fmt.Printf("User %s registered.\n", msg.Username)
	default:
		log.Println("Unknown message type received")
	}
}		

// function to simulate our reddit engine for multi-user scenario
func simulateUsers(rootContext *actor.RootContext, enginePID *actor.PID, numUsers int) {

	//Approach:
	// create 10 users (Done)
	// create 5 subreddits
	// for each user, use random function (5 times) and assign the user to rnadom subreedits among available ones
	// run for loop for 20 times, make random posts by users in subreddits
	// run for loop for 20 times, make random comments (non-hierarchial0)
	// upvote 100 times, random posts
	// downvote 100 times random posts
	// make 2 users leave the subreddit
	// get feed for all 5 subreddits

	// creating 10 users
	for i := 0; i < numUsers; i++ {
        username := fmt.Sprintf("user%d", i+1)

        // Register the user.
        rootContext.Send(enginePID, &RegisterUser{Username: username})
	}
}

func main() {

	// a central point for our actors and managing their lifecycle
	system := actor.NewActorSystem() 
	rootContext := system.Root

	// defining the properties of actor anf using it to spawn an instance
	engineProps := actor.PropsFromProducer(func() actor.Actor { return NewEngineActor() })
	enginePID := rootContext.Spawn(engineProps)

    // Simulate 10 users interacting with the system.
    simulateUsers(rootContext, enginePID, 10)

	time.Sleep(5 * time.Second) // Give some time for messages to be processed.
}