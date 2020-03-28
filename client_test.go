package pandascore

import (
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestAccessToken_IsValid(t *testing.T) {
	assert.True(t, AccessToken("sa").IsValid())
	assert.True(t, AccessToken("sample_access_token").IsValid())
	assert.False(t, AccessToken("").IsValid())
}

func TestNew(t *testing.T) {
	_ = os.Setenv(AccessTokenEnvironmentVariable, "env_access_token")
	result := New()

	assert.NotNil(t, result)
	assert.IsType(t, &Client{}, result)
	assert.EqualValues(t, "env_access_token", result.AccessToken)
}

func TestNew_withoutAccessTokenEnvVariable(t *testing.T) {
	_ = os.Unsetenv(AccessTokenEnvironmentVariable)
	result := New()

	assert.NotNil(t, result)
	assert.IsType(t, &Client{}, result)
	assert.EqualValues(t, "", result.AccessToken)
}

func TestNew_withExplicitAccessTokenSetting(t *testing.T) {
	result := New()
	result.AccessToken = "explicit_access_token"

	assert.NotNil(t, result)
	assert.IsType(t, &Client{}, result)
	assert.EqualValues(t, "explicit_access_token", result.AccessToken)
}

func TestRequest(t *testing.T) {
	defer gock.Off()
	defer assert.True(t, gock.IsDone())

	gock.New("https://api.pandascore.co/csgo/series/running").
		Reply(http.StatusOK).
		File("testdata/csgo-series-running.json")

	client := New()
	series := new([]Series)
	err := client.Request(CSGO, "series/running", series)

	assert.Nil(t, err)
	assert.NotNil(t, series)
	assert.Len(t, *series, 2)
}

func TestRequest_invalidGame(t *testing.T) {
	client := New()
	err := client.Request(Game("doesn't exist"), "series/running", nil)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "doesn't exist")
}

func TestRequest_missingAccessToken(t *testing.T) {
	defer gock.Off()
	defer assert.True(t, gock.IsDone())

	gock.New("https://api.pandascore.co/csgo/series/running").
		Reply(http.StatusForbidden).
		File("testdata/error-missing-access-token.json")

	client := New()
	err := client.Request(CSGO, "series/running", nil)

	assert.NotNil(t, err)
	assert.IsType(t, &PandaScoreError{}, err)
	assert.EqualError(t, err, "PandaScore error: Token is missing")
}
