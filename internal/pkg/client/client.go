package client

import (
	"net/http"
)

// Client holds a base URL for HTTP GET calls.
type Client struct {
	url string
}

// NewClient initializes a Client struct.
func NewClient(baseURL string) Client {
	return Client{baseURL}
}

// MakeRequest makes an HTTP GET request for the given endpoint.
func (c Client) MakeRequest(endpoint string) (*http.Response, error) {
	return http.Get(c.url + endpoint)
}
