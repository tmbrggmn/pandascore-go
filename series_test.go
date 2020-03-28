package pandascore

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestGetRunningSeries(t *testing.T) {
	defer gock.Off()
	defer assert.True(t, gock.IsDone())

	gock.New("https://api.pandascore.co/csgo/series/running").
		Reply(http.StatusOK).
		File("testdata/csgo-series-running.json")

	client := New()
	result, err := client.GetRunningSeries(CSGO)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.IsType(t, []Series{}, result)
	assert.Len(t, result, 2)
}
