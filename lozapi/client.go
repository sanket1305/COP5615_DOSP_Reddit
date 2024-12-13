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

type Monster struct {
	Message string `json:"Message"` // Ensure field names match JSON keys
}

func (c *Client) GetMonsters() (*Monster, error) {
	req, err := http.NewRequest("GET", c.baseUrl+"user/register", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create monsters request: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to submit monsters http request: %v", err)
	}
	defer resp.Body.Close() // Ensure the body is closed

	var monster Monster
	if err := json.NewDecoder(resp.Body).Decode(&monster); err != nil {
		return nil, fmt.Errorf("failed to unmarshal monsters http response: %v", err)
	}

	return &monster, nil
}
