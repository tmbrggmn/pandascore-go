package pandascore

import "time"

// Returns all known leagues for the given game.
func (c *Client) GetAllLeagues(game Game) ([]League, error) {
	leagues := new([]League)
	_, err := c.Request(game, "leagues").GetAll(leagues)
	return *leagues, err
}

// League represents a logical group of series, which are events that belong to a league.
//
// More information: https://developers.pandascore.co/doc/#section/Introduction/Events-hierarchy
type League struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Modified time.Time `json:"modified_at"`
	URL      string    `json:"url"`
}
