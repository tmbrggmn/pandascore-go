package pandascore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequest_Filter(t *testing.T) {
	request := new(Request).Filter("field", "value")

	assert.NotNil(t, request)
	assert.Len(t, request.filter, 1)
	assert.Equal(t, map[string]string{"field": "value"}, request.filter)
}
