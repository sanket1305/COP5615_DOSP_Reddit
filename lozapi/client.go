package lozapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const BaseUrl = "http://localhost:8080/"

type Client struct {
	baseUrl    string
	httpClient *http.Client
}

func NewClient(baseUrl string, httpClient *http.Client) *Client {
	return &Client{
		baseUrl:    baseUrl,
		httpClient: httpClient,
	}
}

type DisplayMessages struct {
	User string
	Conversation map[string][][]string
}

type PostinSubreddit struct{
	PostID			string
	Content			string
	UserID			string
	SubredditName	string
	upvotes			int
	downvotes		int
}

type ResponsePosts struct {
	Posts []PostinSubreddit
}

type Arr struct {
	Arr []string
}

type Message struct {
	Message string `json:"Message"` // Ensure field names match JSON keys
}

func (c *Client) RegisterUser() (*Message, error) {
	req, err := http.NewRequest("GET", c.baseUrl+"user/register", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create reigtser user request: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to submit register user http request: %v", err)
	}
	defer resp.Body.Close() // Ensure the body is closed

	var message Message
	if err := json.NewDecoder(resp.Body).Decode(&message); err != nil {
		return nil, fmt.Errorf("failed to read http response: %v", err)
	}

	return &message, nil
}

func (c *Client) CreateSubreddit(subredditName string) (*Message, error) {
	req, err := http.NewRequest("GET", c.baseUrl+"subreddit/create", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create subreddit request: %v", err)
	}

	reqUrl := req.URL
	queryParams := req.URL.Query()
	queryParams.Set("subredditname", subredditName)
	reqUrl.RawQuery = queryParams.Encode()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to submit subreddit http request: %v", err)
	}
	defer resp.Body.Close() // Ensure the body is closed

	var message Message
	if err := json.NewDecoder(resp.Body).Decode(&message); err != nil {
		return nil, fmt.Errorf("failed to read http response: %v", err)
	}

	return &message, nil
}


func (c *Client) JoinSubreddit(userName string, subredditName string) (*Message, error) {
	req, err := http.NewRequest("GET", c.baseUrl+"subreddit/join", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create subreddit request: %v", err)
	}

	reqUrl := req.URL
	queryParams := req.URL.Query()
	queryParams.Set("subredditname", subredditName)
	queryParams.Set("username", userName)
	reqUrl.RawQuery = queryParams.Encode()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to submit subreddit http request: %v", err)
	}
	defer resp.Body.Close() // Ensure the body is closed

	var message Message
	if err := json.NewDecoder(resp.Body).Decode(&message); err != nil {
		return nil, fmt.Errorf("failed to read http response: %v", err)
	}

	return &message, nil
}


func (c *Client) LeaveSubreddit(userName string, subredditName string) (*Message, error) {
	req, err := http.NewRequest("GET", c.baseUrl+"subreddit/leave", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create subreddit request: %v", err)
	}

	reqUrl := req.URL
	queryParams := req.URL.Query()
	queryParams.Set("subredditname", subredditName)
	queryParams.Set("username", userName)
	reqUrl.RawQuery = queryParams.Encode()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to submit subreddit http request: %v", err)
	}
	defer resp.Body.Close() // Ensure the body is closed

	var message Message
	if err := json.NewDecoder(resp.Body).Decode(&message); err != nil {
		return nil, fmt.Errorf("failed to read http response: %v", err)
	}

	return &message, nil
}


func (c *Client) PostInSubreddit(userName string, subredditName string, content string) (*Message, error) {
	req, err := http.NewRequest("GET", c.baseUrl+"subreddit/post", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create subreddit request: %v", err)
	}

	reqUrl := req.URL
	queryParams := req.URL.Query()
	queryParams.Set("subredditname", subredditName)
	queryParams.Set("username", userName)
	queryParams.Set("content", content)
	reqUrl.RawQuery = queryParams.Encode()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to submit subreddit http request: %v", err)
	}
	defer resp.Body.Close() // Ensure the body is closed

	var message Message
	if err := json.NewDecoder(resp.Body).Decode(&message); err != nil {
		return nil, fmt.Errorf("failed to read http response: %v", err)
	}

	return &message, nil
}


func (c *Client) CommentInSubreddit(userName string, subredditName string, post string, comment string) (*Message, error) {
	req, err := http.NewRequest("GET", c.baseUrl+"subreddit/post/comment", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create subreddit request: %v", err)
	}

	reqUrl := req.URL
	queryParams := req.URL.Query()
	queryParams.Set("subredditname", subredditName)
	queryParams.Set("username", userName)
	queryParams.Set("post", post)
	queryParams.Set("comment", comment)
	reqUrl.RawQuery = queryParams.Encode()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to submit subreddit http request: %v", err)
	}
	defer resp.Body.Close() // Ensure the body is closed

	var message Message
	if err := json.NewDecoder(resp.Body).Decode(&message); err != nil {
		return nil, fmt.Errorf("failed to read http response: %v", err)
	}

	return &message, nil
}


func (c *Client) GetFeed(subredditName string) error {
	req, err := http.NewRequest("GET", c.baseUrl+"subreddit/feed", nil)
	if err != nil {
		return fmt.Errorf("failed to create subreddit request: %v", err)
	}

	reqUrl := req.URL
	queryParams := req.URL.Query()
	queryParams.Set("subredditname", subredditName)
	reqUrl.RawQuery = queryParams.Encode()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to submit subreddit http request: %v", err)
	}
	defer resp.Body.Close() // Ensure the body is closed

	// as different type of responses are expected
	// first get raaw json
	var rawResponse json.RawMessage
	if err := json.NewDecoder(resp.Body).Decode(&rawResponse); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	// now first check if repsonse type is ResponsePosts
	var responsePosts ResponsePosts
	if err := json.Unmarshal(rawResponse, &responsePosts); err == nil {
		fmt.Println("Received Posts:")
		for _, post := range responsePosts.Posts {
			fmt.Printf("ID: %s, Content: %s\n", post.PostID, post.Content)
		}
		return nil
	}

	var errorResponse Message
	if err := json.Unmarshal(rawResponse, &errorResponse); err == nil && errorResponse.Message != "" {
		fmt.Println("Received error message:", errorResponse.Message)
		return nil
	}

	// var arr PostArr
	// if err := json.NewDecoder(resp.Body).Decode(&arr); err != nil {
	// 	return fmt.Errorf("failed to read http response: %v", err)
	// }

	// return &arr, nil
	return nil
}


func (c *Client) GetListOfAvailableSubreddits() (*Arr, error) {
	req, err := http.NewRequest("GET", c.baseUrl+"subreddit/list", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create subreddit request: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to submit subreddit http request: %v", err)
	}
	defer resp.Body.Close() // Ensure the body is closed

	var arr Arr
	if err := json.NewDecoder(resp.Body).Decode(&arr); err != nil {
		return nil, fmt.Errorf("failed to read http response: %v", err)
	}

	return &arr, nil
}


func (c *Client) GetListOfAvailableUsers() (*Arr, error) {
	req, err := http.NewRequest("GET", c.baseUrl+"user/list", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create subreddit request: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to submit subreddit http request: %v", err)
	}
	defer resp.Body.Close() // Ensure the body is closed

	var arr Arr
	if err := json.NewDecoder(resp.Body).Decode(&arr); err != nil {
		return nil, fmt.Errorf("failed to read http response: %v", err)
	}

	return &arr, nil
}


func (c *Client) CheckInbox(username string) (*DisplayMessages, error) {
	req, err := http.NewRequest("GET", c.baseUrl+"user/inbox", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create subreddit request: %v", err)
	}

	reqUrl := req.URL
	queryParams := req.URL.Query()
	queryParams.Set("username", username)
	reqUrl.RawQuery = queryParams.Encode()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to submit subreddit http request: %v", err)
	}
	defer resp.Body.Close() // Ensure the body is closed

	var disp DisplayMessages
	if err := json.NewDecoder(resp.Body).Decode(&disp); err != nil {
		return nil, fmt.Errorf("failed to read http response: %v", err)
	}

	return &disp, nil
}


func (c *Client) SendMessage(sender string, receiver string, message string) (*Message, error) {
	req, err := http.NewRequest("GET", c.baseUrl+"user/sendmessage", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create subreddit request: %v", err)
	}

	reqUrl := req.URL
	queryParams := req.URL.Query()
	queryParams.Set("sender", sender)
	queryParams.Set("receiver", receiver)
	queryParams.Set("message", message)
	reqUrl.RawQuery = queryParams.Encode()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to submit subreddit http request: %v", err)
	}
	defer resp.Body.Close() // Ensure the body is closed

	var msg Message
	if err := json.NewDecoder(resp.Body).Decode(&msg); err != nil {
		return nil, fmt.Errorf("failed to read http response: %v", err)
	}

	return &msg, nil
}

