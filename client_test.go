package pandascore

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
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
