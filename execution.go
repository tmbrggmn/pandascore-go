package pandascore

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Execute the request against the PandaScore API and marshal the response body in the value pointed to by value. In
// case there was an error executing the request, an empty response struct is returned.
//
// Important note: state of the request is maintained ever after execution, so be sure to update any sort or paging
// parameters after execution if you need to.
func (r *Request) Get() (Response, error) {
	if !r.game.IsValid() {
		return Response{}, fmt.Errorf("unknown game '%s'", r.game)
	}

	request, err := buildRequest(r)
	if err != nil {
		log.Printf("unable to build new PandaScore request: %s", err)
		return Response{}, err
	}

	httpResponse, err := r.client.httpClient.Do(request)
	if err != nil {
		log.Printf("PandaScore request failed with error: %s", err)
		return Response{}, err
	}

	err = unmarshallResponseBody(httpResponse, r.value)
	if err != nil {
		log.Printf("failed to unmarshal PandaScore response: %s", err)
		return Response{}, err
	}

	return constructResponse(httpResponse), nil
}

func constructResponse(httpResponse *http.Response) Response {
	getHeaderOrInt := func(header string, defaultValue int) int {
		if result, err := strconv.Atoi(httpResponse.Header.Get(header)); err == nil {
			return result
		} else {
			return defaultValue
		}
	}

	return Response{
		CurrentPage:    getHeaderOrInt("X-Page", 0),
		ResultsPerPage: getHeaderOrInt("X-Per-Page", 0),
		TotalResults:   getHeaderOrInt("X-Total", 0),
	}
}

func buildRequest(request *Request) (*http.Request, error) {
	requestURL := request.client.baseURL.ResolveReference(&url.URL{Path: string(request.game) + "/" + request.path})
	requestURL.RawQuery = setQueryParameters(request, requestURL.Query())

	// Add the bearer token if it's set in the request
	httpRequest, err := http.NewRequest("GET", requestURL.String(), nil)
	if err != nil {
		return nil, err
	} else {
		setAuthorizationHeader(request, httpRequest)
	}

	return httpRequest, nil
}

func setQueryParameters(request *Request, query url.Values) string {
	addQueryParameterFromMap("filter", request.filter, query)
	addQueryParameterFromMap("search", request.search, query)
	addSortingQueryParameter(request.sort, query)
	addPagingQueryParameter(request.page, query)
	return query.Encode()
}

func addQueryParameterFromMap(parameter string, keysAndValues map[string]string, query url.Values) {
	if keysAndValues != nil && len(keysAndValues) > 0 {
		for key, value := range keysAndValues {
			query.Set(parameter+"["+key+"]", value)
		}
	}
}

func addSortingQueryParameter(sorting []string, query url.Values) {
	if len(sorting) > 0 {
		query.Add("sort", strings.Join(sorting, ","))
	}
}

func addPagingQueryParameter(page int, query url.Values) {
	if page > 1 {
		query.Add("page", strconv.Itoa(page))
	}
}

func setAuthorizationHeader(request *Request, httpRequest *http.Request) {
	if len(request.client.accessToken) > 0 {
		httpRequest.Header.Add("Authorization", "Bearer "+request.client.accessToken)
	} else {
		log.Print("Warning: PandaScore access token hasn't been set, requests may fail")
	}
}

func unmarshallResponseBody(response *http.Response, value interface{}) error {
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
