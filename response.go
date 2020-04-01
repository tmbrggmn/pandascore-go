package pandascore

// Response as parsed from a PandaScore API. Important: this is not the actual response body but rather any other
// information (headers, etc.) that might be useful for a caller to get.
type Response struct {
	CurrentPage    int
	ResultsPerPage int
	TotalResults   int
}

// Returns true if there are more pages with more results.
func (r *Response) HasMore() bool {
	return r.TotalResults-(r.ResultsPerPage*r.CurrentPage) > 0
}
