package lozapi


import (
	"encoding/json"
	"fmt"
	"net/http"
)

const BaseUrl = "http://localhost:8080/"

type Client struct {
	baseUrl		string			// base url remains the same, routing apis addresses becomes different based on each request
	httpClient	*http.Client	// standard library in go lang
}

func NewClient(baseUrl string, httpClient *http.Client) *Client {
	return &Client{
		baseUrl: baseUrl,
		httpClient: httpClient,
	}
}

type Monster struct {
	Message			string
}

type GetMosterResponse struct {
	Data Monster
}

func (c *Client) GetMosters() (*GetMosterResponse, error) {
	req, err := http.NewRequest("GET", c.baseUrl + "user/register", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create mosters request: %m", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to submit mosters http request: %m", err)
	}

	var response *GetMosterResponse
	fmt.Printf("%m", resp.Body)
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal monsters http response: %m", err)
	}

	return response, nil
}