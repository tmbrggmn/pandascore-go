package pandascore

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func ExampleGetRunningSeries() {
	defer gock.Off()

	gock.New("https://api.pandascore.co/csgo/series/running").
		Reply(http.StatusOK).
		File("testdata/runningCSGOSeries.json")

	client := New()
	result, err := client.GetRunningSeries(CSGO)

	if err != nil {
		fmt.Printf("Failed to get running series: %s", err)
	} else {
		fmt.Printf("Found %d series with ID %d and %d", len(result), result[0].ID, result[1].ID)
	}

	// Output:
	// Found 2 series with ID 2522 and 2528
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
