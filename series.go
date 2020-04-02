package pandascore

import "time"

// Returns the currently ongoing series for the given game.
func (c *Client) GetRunningSeries(game Game) ([]Series, error) {
	series := new([]Series)

	response, err := c.Request(game, "series/running", series).Get()
	if err != nil {
		return []Series{}, err
	}

	// TODO: this is most likely not the best place to handle paging, it should probably be moved down to request execution
	if response.HasMore() {
		for {
			nextSeries := new([]Series)
			nextPage := response.CurrentPage + 1
			response, err = c.Request(game, "series/running", nextSeries).Page(nextPage).Get()
			if err != nil {
				return nil, err
			}

			*series = append(*series, *nextSeries...)

			if !response.HasMore() {
				break
			}
		}
	}

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
