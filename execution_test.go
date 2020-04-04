package pandascore

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestRequest_Get(t *testing.T) {
	defer gock.Off()
	defer assert.True(t, gock.IsDone())

	gock.New("https://api.pandascore.co/csgo/series/running").
		MatchHeader("Authorization", "test_access_token").
		Reply(http.StatusOK).
		File("testdata/csgo-series-running.json").
		SetHeader("X-Page", "1").
		SetHeader("X-Per-Page", "20").
		SetHeader("X-Total", "20")

	series := new([]Series)
	response, err := New().
		AccessToken("test_access_token").
		Request(CSGO, "series/running").
		Get(series)

	assert.Nil(t, err)
	assert.NotNil(t, series)
	assert.Len(t, *series, 2)
	assert.NotNil(t, response)
	assert.Equal(t, Response{CurrentPage: 1, ResultsPerPage: 20, TotalResults: 20}, response)
	assert.False(t, response.HasMore())
}

func TestRequest_GetAll(t *testing.T) {
	defer gock.Off()
	defer assert.True(t, gock.IsDone())

	gock.New("https://api.pandascore.co/csgo/series/running").
		Reply(http.StatusOK).
		File("testdata/csgo-series-running.json").
		SetHeader("X-Page", "1").
		SetHeader("X-Per-Page", "20").
		SetHeader("X-Total", "40")

	gock.New("https://api.pandascore.co/csgo/series/running").
		MatchParam("page", "2").
		Reply(http.StatusOK).
		File("testdata/csgo-series-running2.json").
		SetHeader("X-Page", "2").
		SetHeader("X-Per-Page", "20").
		SetHeader("X-Total", "40")

	seriesPtr := new([]Series)
	response, err := New().
		AccessToken("test_access_token").
		Request(CSGO, "series/running").
		GetAll(seriesPtr)

	series := *seriesPtr

	assert.Nil(t, err)
	assert.NotNil(t, series)
	assert.Len(t, series, 4)
	assert.Equal(t, []int{2522, 2528, 2523, 2529}, []int{series[0].ID, series[1].ID, series[2].ID, series[3].ID})
	assert.NotNil(t, response)
	assert.Equal(t, Response{CurrentPage: 2, ResultsPerPage: 20, TotalResults: 40}, response)
	assert.False(t, response.HasMore())
}

func TestRequest_Get_invalidGame(t *testing.T) {
	_, err := New().Request(Game("doesn't exist"), "series/running").Get(nil)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "doesn't exist")
}

func TestRequest_Get_missingAccessToken(t *testing.T) {
	defer gock.Off()
	defer assert.True(t, gock.IsDone())

	gock.New("https://api.pandascore.co/csgo/series/running").
		Reply(http.StatusForbidden).
		File("testdata/error-missing-access-token.json")

	_, err := New().Request(CSGO, "series/running").Get(nil)

	assert.NotNil(t, err)
	assert.IsType(t, &PandaScoreError{}, err)
	assert.EqualError(t, err, "PandaScore error: Token is missing")
}

func TestRequest_Get_Filter(t *testing.T) {
	defer gock.Off()
	defer assert.True(t, gock.IsDone())

	gock.New("https://api.pandascore.co/csgo/leagues").
		MatchParam("filter[name]", "ESL").
		MatchParam("filter[slug]", "cs-go-esl").
		Reply(http.StatusOK).
		File("testdata/csgo-leagues-esl.json")

	leagues := new([]League)
	_, err := New().
		Request(CSGO, "leagues").
		Filter("name", "ESL").
		Filter("slug", "cs-go-esl").
		Get(leagues)

	assert.Nil(t, err)
	assert.NotNil(t, leagues)
	assert.Len(t, *leagues, 1)
}

func TestRequest_Get_Filter_withMultipleValues(t *testing.T) {
	defer gock.Off()
	defer assert.True(t, gock.IsDone())

	gock.New("https://api.pandascore.co/csgo/leagues").
		MatchParam("filter[name]", "ESL,IEM").
		MatchParam("filter[slug]", "cs-go-esl").
		Reply(http.StatusOK).
		File("testdata/csgo-leagues-esl.json")

	leagues := new([]League)
	_, err := New().
		Request(CSGO, "leagues").
		Filter("name", "ESL", "IEM").
		Filter("slug", "cs-go-esl").
		Get(leagues)

	assert.Nil(t, err)
	assert.NotNil(t, leagues)
	assert.Len(t, *leagues, 1)
}

func TestRequest_Get_Search(t *testing.T) {
	defer gock.Off()
	defer assert.True(t, gock.IsDone())

	gock.New("https://api.pandascore.co/csgo/leagues").
		MatchParam("search[name]", "ESL").
		MatchParam("search[slug]", "cs-go-esl").
		Reply(http.StatusOK).
		File("testdata/csgo-leagues-esl.json")

	leagues := new([]League)
	_, err := New().
		Request(CSGO, "leagues").
		Search("name", "ESL").
		Search("slug", "cs-go-esl").
		Get(leagues)

	assert.Nil(t, err)
	assert.NotNil(t, leagues)
	assert.Len(t, *leagues, 1)
}

func TestRequest_Get_Sort(t *testing.T) {
	defer gock.Off()
	defer assert.True(t, gock.IsDone())

	gock.New("https://api.pandascore.co/csgo/leagues").
		MatchParam("sort", "name,-modified_at").
		Reply(http.StatusOK).
		File("testdata/csgo-leagues-esl.json")

	leagues := new([]League)
	_, err := New().
		Request(CSGO, "leagues").
		Sort("name", Ascending).
		Sort("modified_at", Descending).
		Get(leagues)

	assert.Nil(t, err)
	assert.NotNil(t, leagues)
	assert.Len(t, *leagues, 1)
}

func Test_constructResponse(t *testing.T) {
	header := http.Header{}
	header.Add("X-Page", "1")
	header.Add("X-Per-Page", "2")
	header.Add("X-Total", "3")
	httpResponse := http.Response{Header: header}
	result := constructResponse(&httpResponse)

	assert.NotNil(t, result)
	assert.Equal(t, Response{CurrentPage: 1, ResultsPerPage: 2, TotalResults: 3}, result)
}
