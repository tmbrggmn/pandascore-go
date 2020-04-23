package pandascore

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestClient_GetAllUpcomingMatchesForSeries(t *testing.T) {
	defer gock.Off()
	defer assert.True(t, gock.IsDone())

	gock.New("https://api.pandascore.co/csgo/matches/upcoming").
		MatchParam("filter[serie_id]", strconv.Itoa(2528)).
		Reply(http.StatusOK).
		File("testdata/csgo-matches-upcoming.json")

	client := New()
	result, err := client.GetAllUpcomingMatchesForSeries(CSGO, 2528)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.IsType(t, []Match{}, result)
	assert.Len(t, result, 3)
}

func TestClient_GetAllUpcomingMatches(t *testing.T) {
	defer gock.Off()
	defer assert.True(t, gock.IsDone())

	gock.New("https://api.pandascore.co/csgo/matches/upcoming").
		MatchParam("page[size]", strconv.Itoa(100)).
		Reply(http.StatusOK).
		File("testdata/csgo-matches-upcoming.json")

	client := New()
	result, err := client.GetAllUpcomingMatches(CSGO)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.IsType(t, []Match{}, result)
	assert.Len(t, result, 3)
	assert.NotNil(t, result[0].League)
	assert.NotNil(t, result[0].Series)
	assert.NotNil(t, result[0].Videogame)
}
