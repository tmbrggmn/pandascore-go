package pandascore

import (
	"net/http"
	"strconv"
	"testing"
	"time"

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
}

func TestClient_GetAllRunningMatches(t *testing.T) {
	defer gock.Off()
	defer assert.True(t, gock.IsDone())

	gock.New("https://api.pandascore.co/csgo/matches/running").
		MatchParam("page[size]", strconv.Itoa(100)).
		Reply(http.StatusOK).
		File("testdata/csgo-matches-running.json")

	client := New()
	result, err := client.GetAllRunningMatches(CSGO)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.IsType(t, []Match{}, result)
	assert.Len(t, result, 4)
	assert.Equal(t, "https://www.twitch.tv/esl_csgo", result[0].LiveURL)
	assert.Equal(t, "https://cdn.pandascore.co/images/team/image/3212/FAZE_CLAN.png", result[0].Opponents[0].Opponent.LogoURL)
}

func TestClient_GetAllUpcomingMatchesBetween(t *testing.T) {
	defer gock.Off()
	defer assert.True(t, gock.IsDone())

	now := time.Now()

	gock.New("https://api.pandascore.co/csgo/matches/upcoming").
		MatchParam("page[size]", strconv.Itoa(100)).
		MatchParam("range[begin_at]", now.UTC().Format(time.RFC3339)+","+now.UTC().Add(time.Hour*24).Format(time.RFC3339)).
		Reply(http.StatusOK).
		File("testdata/csgo-matches-upcoming.json")

	client := New()
	result, err := client.GetAllUpcomingMatchesBetween(CSGO, now, now.Add(time.Hour*24))

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.IsType(t, []Match{}, result)
	assert.Len(t, result, 3)
}
