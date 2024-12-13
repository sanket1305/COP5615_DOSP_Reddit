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


func (c *Client) GetFeed(subredditName string) (*Message, error) {
	req, err := http.NewRequest("GET", c.baseUrl+"subreddit/feed", nil)
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

