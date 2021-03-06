package pandascore

import "time"

// Returns all currently ongoing series for the given game.
func (c *Client) GetAllRunningSeries(game Game) ([]Series, error) {
	series := new([]Series)
	_, err := c.Request(game, "series/running").GetAll(series)
	return *series, err
}

// Series represents an instance of a league event.
//
// More information: https://developers.pandascore.co/doc/#section/Introduction/Events-hierarchy
type Series struct {
	ID       int       `json:"id"`
	FullName string    `json:"full_name"`
	Modified time.Time `json:"modified_at"`
}
