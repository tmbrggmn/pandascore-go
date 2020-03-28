package pandascore

import (
	"fmt"
	"log"
)

// Returns the currently ongoing series for the given game.
func (c *Client) GetRunningSeries(game Game) ([]Series, error) {
	if !game.IsValid() {
		return nil, fmt.Errorf("unknown game '%s'", game)
	}

	request, err := c.buildRequest(game, "series/running")
	if err != nil {
		log.Printf("Unable to build new PandaScore request: %s", err)
		return nil, err
	}

	response, err := c.executeRequest(request)
	if err != nil {
		log.Printf("PandaScore request failed with error: %s", err)
		return nil, err
	}

	series := new([]Series)
	err = c.unmarshallResponse(response, series)
	if err != nil {
		log.Printf("Failed to unmarshal PandaScore response: %s", err)
	}
	return *series, err
}

// Series represents an instance of a league event.
//
// See Also
//
// https://developers.pandascore.co/doc/#section/Introduction/Events-hierarchy
type Series struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
}
