// +build integration

package pandascore

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func loadEnvironmentVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Must have .env file with environment variables in project root")
	}
}

func TestIntegration_getLeagues(t *testing.T) {
	loadEnvironmentVariables()

	leagues, err := New().GetLeagues(CSGO)

	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(leagues), 1)
	assert.NotEmpty(t, leagues[0].ID)
	assert.NotEmpty(t, leagues[0].Name)
	assert.NotEmpty(t, leagues[0].Modified)
}
