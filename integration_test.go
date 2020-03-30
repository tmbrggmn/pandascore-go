// +build integration

package pandascore

import (
	"log"
	"os"
	"strconv"
	"testing"

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

func TestIntegration_getCSGOLeagues(t *testing.T) {
	loadEnvironmentVariables()

	leagues, err := New().GetLeagues(CSGO)

	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(leagues), 1)
	assert.NotEmpty(t, leagues[0].ID)
	assert.NotEmpty(t, leagues[0].Name)
	assert.NotEmpty(t, leagues[0].Modified)

	outputLeaguesAsTable(leagues)
}

func outputLeaguesAsTable(leagues []League) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Modified"})
	for _, league := range leagues {
		table.Append([]string{strconv.Itoa(league.ID), league.Name, league.Modified.String()})
	}
	table.Render()
}
