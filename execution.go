package pandascore

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

// Execute a single request against the PandaScore API to fetch the first page of results and marshal the response body
// in the struct pointed to by value. Typically value will be an array of some struct, eg. []Series
//
// In case there was an error executing the request, an empty response struct is returned.
func (r *Request) Get(value interface{}) (Response, error) {
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

	err = unmarshallResponseBody(httpResponse, value)
	if err != nil {
		log.Printf("failed to unmarshal PandaScore response: %s", err)
		return Response{}, err
	}

	return constructResponse(httpResponse), nil
}

// Execute multiple requests against the PandaScore API to fetch all results from all pages and marshal the response
// body in the struct pointed to by value. Typically value will be an array of some struct, eg. []League
//
// In case there was an error executing the request, an empty response struct is returned.
func (r *Request) GetAll(value interface{}) (Response, error) {
	// Get the first page of results and store them into a generic map so we can merge all results into that
	jsonResponseAsMap := new([]map[string]interface{})
	response, err := r.Get(jsonResponseAsMap)
	if err != nil {
		return Response{}, err
	}

	// Keep fetching results as long as there are more pages and merge the next results into the jsonResponseAsMap
	if response.HasMore() {
		for {
			nextJsonResponseAsMap := new([]map[string]interface{})
			nextPage := response.CurrentPage + 1

			response, err = r.Page(nextPage).Get(nextJsonResponseAsMap)
			if err != nil {
				return Response{}, err
			}

			*jsonResponseAsMap = append(*jsonResponseAsMap, *nextJsonResponseAsMap...)

			if !response.HasMore() {
				break
			}
		}
	}

	// Convert the map to JSON, then back to struct
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	mergedJsonMapAsJson, err := json.Marshal(jsonResponseAsMap)
	if err != nil {
		log.Printf("Failed to marshall merged response map to JSON: %s", err)
		return Response{}, err
	}
	err = json.Unmarshal(mergedJsonMapAsJson, value)
	if err != nil {
		log.Printf("Failed to unmarshall merged response map to struct: %s", err)
		return Response{}, err
	}

	return response, nil
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
	addQueryParameterFromMap("range", request.ranges, query)
	addSortingQueryParameter(request.sort, query)
	addPagingQueryParameter(request.page, query)
	addPageSizeQueryParameter(request.pageSize, query)
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
		query.Add("page[number]", strconv.Itoa(page))
	}
}

func addPageSizeQueryParameter(pageSize int, query url.Values) {
	if pageSize > 0 {
		query.Add("page[size]", strconv.Itoa(pageSize))
	}
}

func setAuthorizationHeader(request *Request, httpRequest *http.Request) {
	if len(request.client.accessToken) > 0 {
		httpRequest.Header.Add("Authorization", "Bearer "+request.client.accessToken)
	} else {
		log.Print("⚠ warning: PandaScore access token hasn't been set, requests may fail")
	}
}

func unmarshallResponseBody(response *http.Response, value interface{}) error {
	defer response.Body.Close()

	// If the response is successful; unmarshal the body. If not; unmarshal the error message in the body.
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	if response.StatusCode >= 200 && response.StatusCode <= 299 {

		return json.NewDecoder(response.Body).Decode(value)
		//return json.NewDecoder(response.Body).Decode(value)
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
