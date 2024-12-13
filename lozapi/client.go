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

	var monster Message
	if err := json.NewDecoder(resp.Body).Decode(&monster); err != nil {
		return nil, fmt.Errorf("failed to read http response: %v", err)
	}

	return &monster, nil
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

	var monster Message
	if err := json.NewDecoder(resp.Body).Decode(&monster); err != nil {
		return nil, fmt.Errorf("failed to read http response: %v", err)
	}

	return &monster, nil
}
