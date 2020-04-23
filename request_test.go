package pandascore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequest_Filter(t *testing.T) {
	request := new(Request).Filter("name", "ESL").Filter("slug", "cs-go-esl")

	assert.NotNil(t, request)
	assert.Len(t, request.filter, 2)
	assert.Equal(t, map[string]string{"name": "ESL", "slug": "cs-go-esl"}, request.filter)
}

func TestRequest_Search(t *testing.T) {
	request := new(Request).
		Search("name", "ESL").
		Search("slug", "cs-go-esl")

	assert.NotNil(t, request)
	assert.Len(t, request.search, 2)
	assert.Equal(t, map[string]string{"name": "ESL", "slug": "cs-go-esl"}, request.search)
}

func TestRequest_Range(t *testing.T) {
	request := new(Request).
		Range("begin_at", "2020-04-23", "2020-04-24").
		Range("modified_at", "2020-04-20", "2020-04-21")

	assert.NotNil(t, request)
	assert.Len(t, request.ranges, 2)
	assert.Equal(t, map[string]string{"begin_at": "2020-04-23,2020-04-24", "modified_at": "2020-04-20,2020-04-21"}, request.ranges)
}

func TestRequest_Sort(t *testing.T) {
	request := new(Request).Sort("name", Descending).Sort("modified_at", Ascending)

	assert.NotNil(t, request)
	assert.Len(t, request.sort, 2)
	assert.Equal(t, []string{"-name", "modified_at"}, request.sort)
}

func TestRequest_Page(t *testing.T) {
	request := new(Request).Page(2)
	assert.Equal(t, 2, request.page, "Expected page to be 2 after it is set to 2")

	request.Page(-1)
	assert.Equal(t, 1, request.page, "Expected page to be 1 after it is set to a negative value")
}

func TestRequest_PageSize(t *testing.T) {
	request := new(Request).PageSize(75)
	assert.Equal(t, 75, request.pageSize, "Expected page to be 75 after it is set to 75")

	request.PageSize(-1)
	assert.Equal(t, 50, request.pageSize, "Expected page size to be 50 after it is set to a negative value")
}
