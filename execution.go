package pandascore

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// Execute the request against the PandaScore API and marshal the response body in the value pointed to by value.
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
	requestURL := request.client.baseURL.ResolveReference(&url.URL{Path: string(request.game) + "/" + request.path})
	addQueryParameterFromMap(request.filter, "filter", requestURL)
	addQueryParameterFromMap(request.search, "search", requestURL)
	addSortQueryParameter(request, requestURL)

	httpRequest, err := http.NewRequest("GET", requestURL.String(), nil)
	if len(request.client.accessToken) > 0 {
		httpRequest.Header.Add("Authorization", "Bearer "+request.client.accessToken)
	}
	return httpRequest, err
}

func addQueryParameterFromMap(keysAndValues map[string]string, parameter string, requestURL *url.URL) {
	if keysAndValues != nil && len(keysAndValues) > 0 {
		query := requestURL.Query()
		for key, value := range keysAndValues {
			query.Set(parameter+"["+key+"]", value)
		}
		requestURL.RawQuery = query.Encode()
	}
}

func addSortQueryParameter(request *Request, requestURL *url.URL) {
	query := requestURL.Query()
	query.Add("sort", strings.Join(request.sort, ","))
	requestURL.RawQuery = query.Encode()
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
