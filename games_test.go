package pandascore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGame_IsValid(t *testing.T) {
	assert.True(t, CSGO.IsValid())
	assert.True(t, Dota2.IsValid())
	assert.True(t, LoL.IsValid())
	assert.False(t, Game("doesn't exist").IsValid())
}
