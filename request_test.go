package pandascore

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestRequest_Execute(t *testing.T) {
	defer gock.Off()
	defer assert.True(t, gock.IsDone())

	gock.New("https://api.pandascore.co/csgo/series/running").
		Reply(http.StatusOK).
		File("testdata/csgo-series-running.json")

	series := new([]Series)
	err := New().Request(CSGO, "series/running", series).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, series)
	assert.Len(t, *series, 2)
}

func TestRequest_Execute_invalidGame(t *testing.T) {
	err := New().Request(Game("doesn't exist"), "series/running", nil).Execute()

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "doesn't exist")
}

func TestRequest_Execute_missingAccessToken(t *testing.T) {
	defer gock.Off()
	defer assert.True(t, gock.IsDone())

	gock.New("https://api.pandascore.co/csgo/series/running").
		Reply(http.StatusForbidden).
		File("testdata/error-missing-access-token.json")

	err := New().Request(CSGO, "series/running", nil).Execute()

	assert.NotNil(t, err)
	assert.IsType(t, &PandaScoreError{}, err)
	assert.EqualError(t, err, "PandaScore error: Token is missing")
}
