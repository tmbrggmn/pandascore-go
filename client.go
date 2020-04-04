// Consume the PandaScore API in Go.
//
// TODO look at errors and how they're returned; we should probably be using Wrap() and stuff but I don't fully understand how that works
package pandascore

import (
	"net/http"
	"net/url"
	"os"
)

const (
	// PandaScore API base URL
	BaseURL string = "api.pandascore.co"

	// Name of the environment variable that contains the PandaScore access token
	AccessTokenEnvironmentVariable string = "PANDASCORE_ACCESS_TOKEN"
)

// PandaScore client which is the primary entity.
type Client struct {
	baseURL     *url.URL
	httpClient  *http.Client
	accessToken string
	filter      string
}

// Construct a new PandaScore client with the default URL.
//
// By default, the PandaScore access token will be read from the environment variable defined in the
// AccessTokenEnvironmentVariable constant. You can set it afterwards, for example if there is not environment variable.
func New() *Client {
	c := &Client{
		httpClient:  http.DefaultClient,
		baseURL:     &url.URL{Scheme: "https", Host: BaseURL},
		accessToken: os.Getenv(AccessTokenEnvironmentVariable),
	}

	return c
}

// Sets this client's PandaScore access token to the given value.
func (c *Client) AccessToken(accessToken string) *Client {
	c.accessToken = accessToken
	return c
}

// Construct a new request for the given game with the given path.
func (c *Client) Request(game Game, path string) *Request {
	return &Request{
		client: c,
		game:   game,
		path:   path,
	}
}
