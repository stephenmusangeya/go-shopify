package admin

import (
	"fmt"
	"net/http"
)

// Client represents a client to interact with the Shopify Admin API.
type Client struct {
	token   string
	store   string
	client  *http.Client
	baseUrl string
}

// NewClient creates a new client to interact with the Shopify Admin API.
func NewClient(store, token string) *Client {
	return &Client{
		token:   token,
		store:   store,
		client:  &http.Client{},
		baseUrl: fmt.Sprintf("https://%s.myshopify.com/admin/api/%s"),
	}
}

// WithHTTPClient allows you to override the default http.Client used by the client.
func (c *Client) WithHTTPClient(client *http.Client) *Client {
	c.client = client
	return c
}
