package pandascore

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// PandaScore API request with it's attributes
//
// TODO add support for search, range and sorting
type Request struct {
	client  *Client
	game    Game
	path    string
	value   interface{}
	filters map[string]string
}

// Registers a new filter to be applied to the request upon execution, where the given field must match the given value.
//
// For example: Filter("name", "ESL") when request leagues would filter all leagues where the league name is "ESL"
func (r *Request) Filter(field string, value string) *Request {
	if r.filters == nil {
		r.filters = make(map[string]string)
	}
	if len(field) > 0 && len(value) > 0 {
		r.filters[field] = value
	}
	return r
}

// Execute a new request against the PandaScore API and marshal the response body in the value pointed to by value
// parameter. You shouldn't really need to use this method since most of the endpoints (eg. series, matches, ...) are
// abstracted by other methods, but its here if you need it.
func (r *Request) Execute() error {
	if !r.game.IsValid() {
		return fmt.Errorf("unknown game '%s'", r.game)
	}

	request, err := buildRequest(r)
	if err != nil {
		log.Printf("unable to build new PandaScore request: %s", err)
		return err
	}

	response, err := r.client.httpClient.Do(request)
	if err != nil {
		log.Printf("PandaScore request failed with error: %s", err)
		return err
	}

	err = unmarshallResponse(response, r.value)
	if err != nil {
		log.Printf("failed to unmarshal PandaScore response: %s", err)
	}
	return err
}

func buildRequest(request *Request) (*http.Request, error) {
	requestURL := request.client.baseUrl.ResolveReference(&url.URL{Path: string(request.game) + "/" + request.path})
	addFiltersToRequestURL(request.filters, requestURL)
	return http.NewRequest("GET", requestURL.String(), nil)
}

func addFiltersToRequestURL(filters map[string]string, requestURL *url.URL) {
	if filters != nil && len(filters) > 0 {
		query := requestURL.Query()
		for key, value := range filters {
			query.Set(key, value)
		}
		requestURL.RawQuery = query.Encode()
	}
}

func unmarshallResponse(response *http.Response, value interface{}) error {
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
