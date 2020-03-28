// Consume the PandaScore API in Go.
package pandascore

// AccessToken represents a PandaScore access token.
type AccessToken string

// Validates that the access token is valid.
func (at AccessToken) IsValid() bool {
	return len(at) > 1
}

// Game represents a single game in the PandaScore API (eg. csgo, dota2, ...)
type Game string

// Validates that the access token is valid.
func (g Game) IsValid() bool {
	return g == CSGO || g == Dota2
}
