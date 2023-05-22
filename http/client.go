package http

import (
	"fmt"
	"net/http"
	"time"
)

// Client represents an HTTP client for making requests.
type Client struct {
	httpClient *http.Client
}

// NewClient creates a new HTTP client.
func NewClient(timeout time.Duration) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}

// Get sends an HTTP GET request to the specified URL and returns the response.
func (c *Client) Get(url string) (*Response, error) {
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to send GET request: %v", err)
	}
	defer resp.Body.Close()

	response := &Response{
		StatusCode: resp.StatusCode,
	}

	// Read the response body if needed
	// Implement your logic to handle the response body according to your requirements

	return response, nil
}
