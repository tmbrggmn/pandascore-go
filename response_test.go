package pandascore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse_HasMore(t *testing.T) {
	response := Response{CurrentPage: 1, ResultsPerPage: 20, TotalResults: 40}
	assert.Equal(t, true, response.HasMore(), "Expected true because there is a 2nd page")

	response.CurrentPage = 2
	assert.Equal(t, false, response.HasMore(), "Expected false because there is no 3rd page")

	response = Response{}
	assert.Equal(t, false, response.HasMore(), "Expected false because all values are 0")
}
