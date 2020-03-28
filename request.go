package pandascore

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// PandaScore API request with it's attributes
type Request struct {
	client *Client
	game   Game
	path   string
	value  interface{}
}

// Execute a new request against the PandaScore API and marshal the response body in the value pointed to by value
// parameter. You shouldn't really need to use this method since most of the endpoints (eg. series, matches, ...) are
// abstracted by other methods, but its here if you need it.
func (r *Request) Execute() error {
	if !r.game.IsValid() {
		return fmt.Errorf("unknown game '%s'", r.game)
	}

	request, err := r.buildRequest()
	if err != nil {
		log.Printf("unable to build new PandaScore request: %s", err)
		return err
	}

	response, err := r.client.httpClient.Do(request)
	if err != nil {
		log.Printf("PandaScore request failed with error: %s", err)
		return err
	}

	err = r.unmarshallResponse(response, r.value)
	if err != nil {
		log.Printf("failed to unmarshal PandaScore response: %s", err)
	}
	return err
}

func (r *Request) buildRequest() (*http.Request, error) {
	requestURL := r.client.baseUrl.ResolveReference(&url.URL{Path: string(r.game) + "/" + r.path}).String()
	return http.NewRequest("GET", requestURL, nil)
}

func (r *Request) unmarshallResponse(response *http.Response, value interface{}) error {
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

// Represents an error coming directly from the PandaScore API (eg. no or invalid access token).
type PandaScoreError struct {
	Message string `json:"error"`
}

func (pse *PandaScoreError) Error() string {
	return "PandaScore error: " + pse.Message
}
