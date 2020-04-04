package pandascore

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestClient_GetAllUpcomingMatches(t *testing.T) {
	defer gock.Off()
	defer assert.True(t, gock.IsDone())

	gock.New("https://api.pandascore.co/csgo/matches/upcoming").
		MatchParam("filter[serie_id]", strconv.Itoa(2528)).
		Reply(http.StatusOK).
		File("testdata/csgo-matches-upcoming.json")

	client := New()
	result, err := client.GetAllUpcomingMatches(CSGO, Series{ID: 2528})

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.IsType(t, []Match{}, result)
	assert.Len(t, result, 3)
}

func TestClient_GetAllUpcomingMatches_withPaging(t *testing.T) {
	defer gock.Off()
	defer assert.True(t, gock.IsDone())

	gock.New("https://api.pandascore.co/csgo/series/running").
		Reply(http.StatusOK).
		AddHeader("X-Page", "1").
		AddHeader("X-Per-Page", "2").
		AddHeader("X-Total", "4").
		File("testdata/csgo-series-running.json")

	gock.New("https://api.pandascore.co/csgo/series/running").
		MatchParam("page", "2").
		Reply(http.StatusOK).
		AddHeader("X-Page", "2").
		AddHeader("X-Per-Page", "2").
		AddHeader("X-Total", "4").
		File("testdata/csgo-series-running2.json")

	client := New()
	result, err := client.GetAllRunningSeries(CSGO)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.IsType(t, []Series{}, result)
	assert.Len(t, result, 4)
	assert.Equal(t, []int{2522, 2528, 2523, 2529}, []int{result[0].ID, result[1].ID, result[2].ID, result[3].ID})
}
