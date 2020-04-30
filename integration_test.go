// +build integration

package pandascore

import (
	"log"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/olekukonko/tablewriter"
	"github.com/stretchr/testify/assert"
)

func loadEnvironmentVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Missing .env file with " + AccessTokenEnvironmentVariable + " variable in project root")
	}
}

func TestIntegration_getAllLeagues(t *testing.T) {
	loadEnvironmentVariables()

	games := []Game{CSGO, Dota2, LoL}
	for _, game := range games {
		leagues, err := New().GetAllLeagues(game)

		assert.Nil(t, err)
		assert.GreaterOrEqual(t, len(leagues), 50)
		assert.NotEmpty(t, leagues[0].ID)
		assert.NotEmpty(t, leagues[0].Name)
		assert.NotEmpty(t, leagues[0].Modified)

		outputLeaguesAsTable(leagues)
	}
}

func outputLeaguesAsTable(leagues []League) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "ID", "Name", "Modified"})
	for index, league := range leagues {
		table.Append([]string{strconv.Itoa(index + 1), strconv.Itoa(league.ID), league.Name, league.Modified.String()})
	}
	table.Render()
}

func TestIntegration_getAllCSGOUpcomingMatches(t *testing.T) {
	loadEnvironmentVariables()

	matches, err := New().GetAllUpcomingMatches(CSGO)

	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(matches), 101)
	assert.NotEmpty(t, matches[0].ID)
	assert.NotEmpty(t, matches[0].Name)
	assert.NotEmpty(t, matches[0].Modified)
	assert.NotNil(t, matches[0].Series)
	assert.NotNil(t, matches[0].League)

	outputMatchesAsTable(matches)
}

func TestIntegration_getAllCSGOUpcomingMatchesBetween(t *testing.T) {
	loadEnvironmentVariables()

	matches, err := New().GetAllUpcomingMatchesBetween(CSGO, time.Now(), time.Now().Add(time.Hour*24))

	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(matches), 10)
	assert.NotEmpty(t, matches[0].ID)
	assert.NotEmpty(t, matches[0].Name)
	assert.NotEmpty(t, matches[0].Modified)
	assert.NotNil(t, matches[0].Series)
	assert.NotNil(t, matches[0].League)

	outputMatchesAsTable(matches)
}

func TestIntegration_getAllCSGORunningMatches(t *testing.T) {
	loadEnvironmentVariables()

	matches, err := New().GetAllRunningMatches(CSGO)

	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(matches), 2)
	assert.NotEmpty(t, matches[0].ID)
	assert.NotEmpty(t, matches[0].Name)
	assert.NotEmpty(t, matches[0].Modified)
	assert.NotNil(t, matches[0].Series)
	assert.NotNil(t, matches[0].League)
	assert.NotNil(t, matches[0].Videogame)

	outputMatchesAsTable(matches)
}

func outputMatchesAsTable(matches []Match) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "ID", "Name", "Starts at", "Modified", "Video Game", "League", "Series"})
	for index, match := range matches {
		table.Append([]string{
			strconv.Itoa(index + 1),
			strconv.Itoa(match.ID),
			match.Name,
			match.BeginsAt.String(),
			match.Modified.String(),
			match.Videogame.Name,
			match.League.Name,
			match.Series.FullName,
		})
	}
	table.Render()
}
