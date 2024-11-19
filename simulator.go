package main

import (
	"fmt"
    "math/rand"
    "sync"
)

type Simulator struct {
    engine *Engine // Reference to the Reddit-like engine.
    users  []*User // List of simulated users.
}

func NewSimulator(engine *Engine) *Simulator {
    return &Simulator{engine: engine}
}

// Simulate users registering and interacting with subreddits.
func (sim *Simulator) Simulate(numUsers int) {
    var wg sync.WaitGroup

    for i := 0; i < numUsers; i++ {
        wg.Add(1)

        go func(userNum int) {
            defer wg.Done()

            username := fmt.Sprintf("user%d", userNum+1)
            user := sim.engine.RegisterUser(username)

            // Simulate joining/creating subreddits and posting.
            sim.simulateActivity(user)

        }(i)
    }

    wg.Wait() // Wait for all goroutines to finish.
}

func (sim *Simulator) simulateActivity(user *User) {

    subreddits := []string{"golang", "technology", "gaming"}

    // Randomly join or create subreddits.
    for _, sr := range subreddits {
        sim.engine.CreateSubreddit(sr)
        sim.engine.JoinSubreddit(user.ID, sr)

        // Simulate posting in the subreddit.
        if rand.Intn(2) == 0 { // Randomize actions.
            sim.engine.PostInSubreddit(user.ID, sr, "This is a test post!")
        }

        // Simulate voting on posts.
        posts := sim.engine.GetFeed(sr)
        for _, post := range posts {
            if rand.Intn(2) == 0 { // Randomize upvote/downvote.
                sim.engine.VoteOnPost(user.ID, post, true)
            } else {
                sim.engine.VoteOnPost(user.ID, post, false)
            }
        }
    }
}

func main() {

    engine := NewEngine()

    simulator := NewSimulator(engine)

    // Simulate 100 users interacting with the engine.
    simulator.Simulate(100)

}