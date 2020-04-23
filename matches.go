package pandascore

import (
	"strconv"
	"time"
)

// Returns all upcoming matches for the given game & series ID.
func (c *Client) GetAllUpcomingMatchesForSeries(game Game, seriesID int) ([]Match, error) {
	matches := new([]Match)
	_, err := c.Request(game, "matches/upcoming").
		Filter("serie_id", strconv.Itoa(seriesID)).
		Get(matches)
	return *matches, err
}

// Returns all upcoming matches for the given game.
func (c *Client) GetAllUpcomingMatches(game Game) ([]Match, error) {
	matches := new([]Match)
	_, err := c.Request(game, "matches/upcoming").PageSize(100).GetAll(matches)
	return *matches, err
}

// Returns all upcoming matches for the given game starting from the given time up until the given time.
//
// Careful: the given times are always set to UTC (Zulu) so timezones are not take into account.
func (c *Client) GetAllUpcomingMatchesBetween(game Game, beginning time.Time, until time.Time) ([]Match, error) {
	matches := new([]Match)
	_, err := c.Request(game, "matches/upcoming").
		Range("begin_at", beginning.UTC().Format(time.RFC3339), until.UTC().Format(time.RFC3339)).
		PageSize(100).
		GetAll(matches)
	return *matches, err
}

// Returns all running matches for the given game.
func (c *Client) GetAllRunningMatches(game Game) ([]Match, error) {
	matches := new([]Match)
	_, err := c.Request(game, "matches/running").PageSize(100).GetAll(matches)
	return *matches, err
}

// Match represents an instance of a single match between 2 opponents (teams or players).
//
// More information: https://developers.pandascore.co/doc/#section/Introduction/Events-hierarchy
type Match struct {
	ID        int             `json:"id"`
	Name      string          `json:"name"`
	BeginsAt  time.Time       `json:"begin_at"`
	Modified  time.Time       `json:"modified_at"`
	Videogame Videogame       `json:"videogame"`
	Opponents []MatchOpponent `json:"opponents"`
	Series    Series          `json:"serie"`
	League    League          `json:"league"`
}

// MatchOpponent represents an opponent as defined for a specific match. Whether the opponent is a team is defined on
// this level, which I find really weird.
type MatchOpponent struct {
	Type     string   `json:"type"`
	Opponent Opponent `json:"opponent"`
}

// Opponent represents a single opponent that partakes in a match.
type Opponent struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

// Videogame represents the type of game that this match is being played in.
type Videogame struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
