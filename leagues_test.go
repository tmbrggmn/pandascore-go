package pandascore

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestGetLeagues(t *testing.T) {
	defer gock.Off()
	defer assert.True(t, gock.IsDone())

	gock.New("https://api.pandascore.co/csgo/leagues").
		Reply(http.StatusOK).
		File("testdata/csgo-leagues.json")

	client := New()
	result, err := client.GetLeagues(CSGO)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.IsType(t, []League{}, result)
	assert.Len(t, result, 50)
	assert.Equal(t,
		League{
			ID:       4351,
			Name:     "HIPFIRED CUP",
			Modified: time.Date(2020, time.March, 26, 10, 5, 6, 0, time.UTC),
			URL:      "https://hipfired.media/hipfired-cup/",
		},
		result[0],
	)
}
