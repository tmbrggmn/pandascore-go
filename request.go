package pandascore

import (
	"strings"
)

const (
	Ascending  Sorting = 0
	Descending Sorting = 1
)

// Sorting used for specifying sorting parameters.
type Sorting byte

func (o Sorting) forField(field string) string {
	if o == Descending {
		return "-" + field
	} else {
		return field
	}
}

// PandaScore API request with it's attributes
// TODO: add support for range
type Request struct {
	client *Client
	game   Game
	path   string
	value  interface{}
	filter map[string]string
	search map[string]string
	sort   []string
}

// Adds a filter parameter to the request, where the given field must match the given value.
//
// More information: https://developers.pandascore.co/doc/index.htm#section/Introduction/Filtering
func (r *Request) Filter(field string, value ...string) *Request {
	if r.filter == nil {
		r.filter = make(map[string]string)
	}
	if len(field) > 0 && len(value) > 0 {
		r.filter[field] = strings.Join(value, ",")
	}
	return r
}

// Adds a search parameter to the request, where the given field must contain the given value.
//
// More information: https://developers.pandascore.co/doc/index.htm#section/Introduction/Search
func (r *Request) Search(field string, value string) *Request {
	if r.search == nil {
		r.search = make(map[string]string)
	}
	if len(field) > 0 && len(value) > 0 {
		r.search[field] = value
	}
	return r
}

// Adds a sort parameter to the request.
//
// More information: https://developers.pandascore.co/doc/index.htm#section/Introduction/Sorting
func (r *Request) Sort(field string, order Sorting) *Request {
	if r.sort == nil {
		r.sort = []string{}
	}
	if len(field) > 0 {
		r.sort = append(r.sort, order.forField(field))
	}
	return r
}
