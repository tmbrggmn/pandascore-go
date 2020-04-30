package pandascore

const (
	CSGO  Game = "csgo"
	Dota2 Game = "dota2"
	LoL   Game = "lol"
)

// Game represents a single game in the PandaScore API (eg. csgo, dota2, ...)
type Game string

// Validates that the access token is valid.
func (g Game) IsValid() bool {
	return g == CSGO || g == Dota2 || g == LoL
}
