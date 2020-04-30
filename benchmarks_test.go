package pandascore

import (
	"net/http"
	"testing"

	"gopkg.in/h2non/gock.v1"
)

// Benchmarks large response
func BenchmarkRequest_GetAll_largeFile(b *testing.B) {
	b.ReportAllocs()

	defer gock.Off()

	for n := 0; n < b.N; n++ {
		gock.New("https://api.pandascore.co/csgo/matches/upcoming").
			Reply(http.StatusOK).
			File("testdata/csgo-matches-upcoming-large.json").
			SetHeader("X-Page", "1").
			SetHeader("X-Per-Page", "200").
			SetHeader("X-Total", "200")
	}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		matchesPtr := new([]Match)
		_, err := New().
			AccessToken("test_access_token").
			Request(CSGO, "matches/upcoming").
			GetAll(matchesPtr)

		matches := *matchesPtr

		if err != nil {
			b.Errorf("Expected request to have succeeded, instead got error: %s", err)
		}
		if len(matches) != 200 {
			b.Errorf("Expected response to have 200 items, instead got %d", len(matches))
		}
	}
}
