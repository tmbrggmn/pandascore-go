// Consume the PandaScore API in Go.
package pandascore

import (
	"encoding/json"
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

	CSGO  Game = "csgo"
	Dota2 Game = "dota2"
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

func (c *Client) buildRequestURL(game Game, path string) *url.URL {
	return c.baseUrl.ResolveReference(&url.URL{Path: string(game) + "/" + path})
}

func (c *Client) buildRequest(game Game, path string) (*http.Request, error) {
	return http.NewRequest("GET", c.buildRequestURL(game, path).String(), nil)
}

func (c *Client) executeRequest(request *http.Request) (*http.Response, error) {
	return c.httpClient.Do(request)
}

func (c *Client) unmarshallResponse(response *http.Response, value interface{}) error {
	defer response.Body.Close()

	// If the response is successful, unmarshal the body. If not, get the error message in the body.
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
