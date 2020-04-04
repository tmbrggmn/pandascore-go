package pandascore

import (
	"strconv"
	"time"
)

// Returns all upcoming matched for the given series ID.
func (c *Client) GetAllUpcomingMatches(game Game, series Series) ([]Match, error) {
	matches := new([]Match)
	_, err := c.Request(game, "matches/upcoming").
		Filter("serie_id", strconv.Itoa(series.ID)).
		Get(matches)
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
	Opponents []MatchOpponent `json:"opponents"`
	Series    Series          `json:"serie"`
}

// MatchOpponent represents an opponent as defined for a specific match. Whether the opponent is a team is defined on
// this level, which I find really weird. As if there are opponents which can suddenly not be a team anymore if they're
// not partaking in a match?
type MatchOpponent struct {
	Type     string   `json:"type"`
	Opponent Opponent `json:"opponent"`
}

// Opponent represents a single opponent that partakes in a match.
type Opponent struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}
