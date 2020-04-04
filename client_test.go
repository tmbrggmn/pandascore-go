package pandascore

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	_ = os.Setenv(AccessTokenEnvironmentVariable, "env_access_token")
	result := New()

	assert.NotNil(t, result)
	assert.IsType(t, &Client{}, result)
	assert.EqualValues(t, "env_access_token", result.accessToken)
}

func TestNew_withoutAccessTokenEnvVariable(t *testing.T) {
	_ = os.Unsetenv(AccessTokenEnvironmentVariable)
	result := New()

	assert.NotNil(t, result)
	assert.IsType(t, &Client{}, result)
	assert.EqualValues(t, "", result.accessToken)
}

func TestClient_AccessToken_withExplicitAccessTokenSetting(t *testing.T) {
	result := New().AccessToken("explicit_access_token")

	assert.NotNil(t, result)
	assert.IsType(t, &Client{}, result)
	assert.EqualValues(t, "explicit_access_token", result.accessToken)
}

func TestClient_Request(t *testing.T) {
	result := New().Request(CSGO, "/path")

	assert.NotNil(t, result)
	assert.IsType(t, &Request{}, result)
	assert.Equal(t, CSGO, result.game)
	assert.Equal(t, "/path", result.path)
}
