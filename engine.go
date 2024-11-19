package main

import (
	"fmt"
	"sync"
)

type User struct {
	ID       string
	Username string
	Karma    int
}

type Post struct {
	ID       string
	UserID   string
	Content  string
	Upvotes  int
	Downvotes int
	Comments []*Comment
}

type Comment struct {
	ID       string
	UserID   string
	Content  string
	Upvotes  int
	Downvotes int
	Replies  []*Comment
}

type Subreddit struct {
	Name    string
	Members map[string]*User // Map of user IDs to Users
	Posts   []*Post          // List of posts in the subreddit
}

type Engine struct {
	mu         sync.Mutex // To manage concurrent access to data structures
	Users      map[string]*User
	Subreddits map[string]*Subreddit
}

func NewEngine() *Engine {
	return &Engine{
		Users:      make(map[string]*User),
		Subreddits: make(map[string]*Subreddit),
	}
}

// Register a new user.
func (e *Engine) RegisterUser(username string) *User {
	e.mu.Lock()
	defer e.mu.Unlock()

	user := &User{ID: username, Username: username}
	e.Users[username] = user
	fmt.Printf("User %s registered.\n", username)
	return user
}

// Create a new subreddit.
func (e *Engine) CreateSubreddit(name string) *Subreddit {
	e.mu.Lock()
	defer e.mu.Unlock()

	subreddit := &Subreddit{Name: name, Members: make(map[string]*User)}
	e.Subreddits[name] = subreddit
	fmt.Printf("Subreddit %s created.\n", name)
	return subreddit
}

// Join a subreddit.
func (e *Engine) JoinSubreddit(userID, subredditName string) {
	e.mu.Lock()
	defer e.mu.Unlock()

	user, ok := e.Users[userID]
	if !ok {
		fmt.Printf("User %s not found.\n", userID)
		return
	}
	subreddit, ok := e.Subreddits[subredditName]
	if !ok {
		fmt.Printf("Subreddit %s not found.\n", subredditName)
		return
	}
	subreddit.Members[userID] = user
	fmt.Printf("User %s joined subreddit %s.\n", userID, subredditName)
}

// Leave a subreddit.
func (e *Engine) LeaveSubreddit(userID, subredditName string) {
	e.mu.Lock()
	defer e.mu.Unlock()

	subreddit, ok := e.Subreddits[subredditName]
	if !ok {
		fmt.Printf("Subreddit %s not found.\n", subredditName)
		return
	}
	delete(subreddit.Members, userID)
	fmt.Printf("User %s left subreddit %s.\n", userID, subredditName)
}

// Post in a subreddit.
func (e *Engine) PostInSubreddit(userID, subredditName, content string) *Post {
	e.mu.Lock()
	defer e.mu.Unlock()

	user, ok := e.Users[userID]
	if !ok {
		fmt.Printf("User %s not found.\n", userID)
		return nil
	}
	subreddit, ok := e.Subreddits[subredditName]
	if !ok {
		fmt.Printf("Subreddit %s not found.\n", subredditName)
		return nil
	}
	post := &Post{ID: fmt.Sprintf("%d", len(subreddit.Posts)+1), UserID: user.ID, Content: content}
	subreddit.Posts = append(subreddit.Posts, post)
	fmt.Printf("User %s posted in subreddit %s: %s\n", user.Username, subreddit.Name, content)
	return post
}

// Comment on a post.
func (e *Engine) CommentOnPost(userID, postID, content string, post *Post) *Comment {
	e.mu.Lock()
	defer e.mu.Unlock()

	comment := &Comment{ID: fmt.Sprintf("%d", len(post.Comments)+1), UserID: userID, Content: content}
	post.Comments = append(post.Comments, comment)
	fmt.Printf("User %s commented on post %s: %s\n", userID, post.ID, content)
	return comment
}

// Upvote or downvote a post.
func (e *Engine) VoteOnPost(userID string, post *Post, upvote bool) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if upvote {
		post.Upvotes++
	} else {
		post.Downvotes++
	}
	fmt.Printf("User %s voted on post %s. Upvotes: %d Downvotes: %d\n", userID, post.ID, post.Upvotes, post.Downvotes)
}

// Get feed of posts from a subreddit.
func (e *Engine) GetFeed(subredditName string) []*Post {
	e.mu.Lock()
	defer e.mu.Unlock()

	subreddit, ok := e.Subreddits[subredditName]
	if !ok {
		fmt.Printf("Subreddit %s not found.\n", subredditName)
		return nil
	}
	return subreddit.Posts // Return the list of posts in the subreddit.
}
