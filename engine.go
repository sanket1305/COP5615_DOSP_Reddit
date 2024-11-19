package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/asynkron/protoactor-go/actor"
)

// Message types
type RegisterUser struct {
	Username string
}
type CreateSubreddit struct {
	Name string
}
type JoinSubreddit struct {
	UserID       string
	SubredditName string
}
type PostInSubreddit struct {
	UserID       string
	SubredditName string
	Content      string
}
type GetFeed struct {
	SubredditName string
}

// User Actor
type UserActor struct {
	ID       string
	Username string
	Karma    int
}

func (state *UserActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *RegisterUser:
		state.Username = msg.Username
		fmt.Printf("User %s registered.\n", state.Username)
	}
}

// Subreddit Actor
type SubredditActor struct {
	Name  string
	Posts []*PostActor // List of posts in the subreddit.
}

func (state *SubredditActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *CreateSubreddit:
		state.Name = msg.Name
		fmt.Printf("Subreddit %s created.\n", state.Name)
	case *PostInSubreddit:
		post := &PostActor{UserID: msg.UserID, Content: msg.Content}
		state.Posts = append(state.Posts, post)
		fmt.Printf("User %s posted in subreddit %s: %s\n", msg.UserID, state.Name, msg.Content)
	case *GetFeed:
		fmt.Printf("Fetching feed for subreddit %s\n", state.Name)
		for _, post := range state.Posts {
			fmt.Printf("Post by %s: %s\n", post.UserID, post.Content)
		}
	}
}

// Post Actor (for simplicity, not making it a full actor here)
type PostActor struct {
	UserID  string
	Content string
}

// Engine Actor (Orchestrator)
type EngineActor struct {
	users      map[string]*actor.PID
	subreddits map[string]*actor.PID
}

func NewEngineActor() *EngineActor {
	return &EngineActor{
		users:      make(map[string]*actor.PID),
		subreddits: make(map[string]*actor.PID),
	}
}

func (state *EngineActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {

	case *RegisterUser:
		userProps := actor.PropsFromProducer(func() actor.Actor { return &UserActor{ID: msg.Username} })
		userPID := ctx.Spawn(userProps)
		state.users[msg.Username] = userPID

	case *CreateSubreddit:
		subredditProps := actor.PropsFromProducer(func() actor.Actor { return &SubredditActor{Name: msg.Name} })
		subredditPID := ctx.Spawn(subredditProps)
		state.subreddits[msg.Name] = subredditPID

	case *JoinSubreddit:
		if subredditPID, ok := state.subreddits[msg.SubredditName]; ok {
			ctx.Send(subredditPID, &JoinSubreddit{UserID: msg.UserID})
			fmt.Printf("User %s joined subreddit %s.\n", msg.UserID, msg.SubredditName)
		} else {
			fmt.Printf("Subreddit %s not found.\n", msg.SubredditName)
		}

	case *PostInSubreddit:
		if subredditPID, ok := state.subreddits[msg.SubredditName]; ok {
			ctx.Send(subredditPID, &PostInSubreddit{UserID: msg.UserID, SubredditName: msg.SubredditName, Content: msg.Content})
			fmt.Printf("User %s posted in subreddit %s.\n", msg.UserID, msg.SubredditName)
		} else {
			fmt.Printf("Subreddit %s not found.\n", msg.SubredditName)
		}

	case *GetFeed:
		if subredditPID, ok := state.subreddits[msg.SubredditName]; ok {
			ctx.Send(subredditPID, &GetFeed{SubredditName: msg.SubredditName})
			fmt.Printf("Fetching feed for subreddit %s.\n", msg.SubredditName)
		} else {
			fmt.Printf("Subreddit %s not found.\n", msg.SubredditName)
		}
	default:
		log.Println("Unknown message type received")
	}
}

// Simulator to simulate multiple users and actions
func simulateUsers(rootContext *actor.RootContext, enginePID *actor.PID, numUsers int) {

	for i := 0; i < numUsers; i++ {
        username := fmt.Sprintf("user%d", i+1)

        // Register the user.
        rootContext.Send(enginePID, &RegisterUser{Username: username})

        // Simulate creating and joining subreddits.
        subredditName := fmt.Sprintf("sub%d", rand.Intn(5)+1) // Randomly choose from 5 subreddits.
        rootContext.Send(enginePID, &CreateSubreddit{Name: subredditName})
        rootContext.Send(enginePID, &JoinSubreddit{UserID: username, SubredditName: subredditName})

        // Simulate posting in the subreddit.
        content := fmt.Sprintf("This is a post by %s in subreddit %s.", username, subredditName)
        rootContext.Send(enginePID, &PostInSubreddit{UserID: username, SubredditName: subredditName, Content: content})

        // Fetch the feed for the subreddit.
        rootContext.Send(enginePID, &GetFeed{SubredditName: subredditName})

        // Introduce a small delay to simulate real-time actions.
        // time.Sleep(time.Millisecond * 500)
    }
}

func main() {

	system := actor.NewActorSystem() // Create an actor system.
	rootContext := system.Root       // Get the root context from the system.

	engineProps := actor.PropsFromProducer(func() actor.Actor { return NewEngineActor() })
	enginePID := rootContext.Spawn(engineProps)

    // Simulate 10 users interacting with the system.
    simulateUsers(rootContext, enginePID, 10)

	time.Sleep(5 * time.Second) // Give some time for messages to be processed.
}