package catena

import "net/http"

// Option configures the Client.
type Option func(*Client)

// WithBaseURL sets the base URL for the API.
func WithBaseURL(url string) Option {
	return func(c *Client) {
		c.baseURL = url
	}
}

// WithHTTPClient sets the HTTP client to use for requests.
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// WithClientID sets the client ID for authentication.
func WithClientID(clientID string) Option {
	return func(c *Client) {
		c.clientID = clientID
	}
}

// WithClientSecret sets the client secret for authentication.
func WithClientSecret(clientSecret string) Option {
	return func(c *Client) {
		c.clientSecret = clientSecret
	}
}
