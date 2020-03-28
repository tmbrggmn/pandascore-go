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
		File("testdata/runningCSGOSeries.json")

	client := New()
	result, err := client.GetRunningSeries(CSGO)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 2)
	assert.Equal(t, 2522, result[0].ID)
	assert.Equal(t, "ANZ Champs: Online Stage season 10 2020", result[0].FullName)
	assert.Equal(t, 2528, result[1].ID)
	assert.Equal(t, "Pro League season 11 2020", result[1].FullName)
}

func TestGetRunningSeries_withInvalidAccessToken(t *testing.T) {
	defer gock.Off()
	defer assert.True(t, gock.IsDone())

	gock.New("https://api.pandascore.co/csgo/series/running").
		Reply(http.StatusForbidden).
		File("testdata/missingAccessToken.json")

	client := New()
	_, err := client.GetRunningSeries(CSGO)

	assert.NotNil(t, err)
	assert.IsType(t, &PandaScoreError{}, err)
	assert.EqualError(t, err, "PandaScore error: Token is missing")
}

func TestGetRunningSeries_withInvalidGame(t *testing.T) {
	client := New()
	result, err := client.GetRunningSeries(Game("doesn't exist"))

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "doesn't exist")
	assert.Empty(t, result)
}
