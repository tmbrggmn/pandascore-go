// Consume the PandaScore API in Go.
//
// TODO Add support for filters, search, range and sorting
package pandascore

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

const (
	// BaseURL is the PandaScore API base URL
	BaseURL string = "api.pandascore.co"

	// AccessTokenEnvironmentVariable contains the name of the environment variable that points to the PandaScore
	// access token
	AccessTokenEnvironmentVariable string = "PANDASCORE_ACCESS_TOKEN"
)

// PandaScore client which is the primary entity.
type Client struct {
	baseUrl     *url.URL
	httpClient  *http.Client
	AccessToken AccessToken
}

// Construct a new PandaScore client with the default URL.
//
// By default, the PandaScore ACCESS TOKEN will be read from the environment variable defined in the
// AccessTokenEnvironmentVariable constant. You can overwrite it afterwards, for example if there is not environment
// variable.
func New() *Client {
	c := &Client{
		httpClient:  http.DefaultClient,
		baseUrl:     &url.URL{Scheme: "https", Host: BaseURL},
		AccessToken: AccessToken(os.Getenv(AccessTokenEnvironmentVariable)),
	}

	return c
}

// Execute a new request against the PandaScore API and marshal the response body in the value pointed to by value
// parameter. You shouldn't really need to use this method since most of the endpoints (eg. series, matches, ...) are
// abstracted by other methods, but its here if you need it.
func (c *Client) Request(game Game, path string, value interface{}) error {
	if !game.IsValid() {
		return fmt.Errorf("unknown game '%s'", game)
	}

	request, err := c.buildRequest(game, path)
	if err != nil {
		log.Printf("unable to build new PandaScore request: %s", err)
		return err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		log.Printf("PandaScore request failed with error: %s", err)
		return err
	}

	err = c.unmarshallResponse(response, value)
	if err != nil {
		log.Printf("Failed to unmarshal PandaScore response: %s", err)
	}
	return err
}

func (c *Client) buildRequest(game Game, path string) (*http.Request, error) {
	requestURL := c.baseUrl.ResolveReference(&url.URL{Path: string(game) + "/" + path}).String()
	return http.NewRequest("GET", requestURL, nil)
}

func (c *Client) unmarshallResponse(response *http.Response, value interface{}) error {
	defer response.Body.Close()

	// If the response is successful; unmarshal the body. If not; unmarshal the error message in the body.
	if response.StatusCode >= 200 && response.StatusCode <= 299 {
		return json.NewDecoder(response.Body).Decode(value)
	} else {
		pandaScoreError := new(PandaScoreError)
		err := json.NewDecoder(response.Body).Decode(pandaScoreError)
		if err == nil {
			err = pandaScoreError
		}
		return err
	}
}

// AccessToken represents a PandaScore access token.
// TODO Remove AccessToken type and replace it with a string, it makes things more abstract for no reason
type AccessToken string

// Validates that the access token is valid.
func (at AccessToken) IsValid() bool {
	return len(at) > 1
}

// Represents an error coming directly from the PandaScore API (eg. no or invalid access token).
type PandaScoreError struct {
	Message string `json:"error"`
}

func (pse *PandaScoreError) Error() string {
	return "PandaScore error: " + pse.Message
}
