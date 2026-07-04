package game

import	"math/rand"

// Difficulty represents the bot AI level.
type Difficulty int

// Difficulty levels for bot AI.
const (
	Easy Difficulty = iota
	Medium
	Hard
	Impossible
)

// DifficultyConfig controls bot behaviour.
type DifficultyConfig struct {
	MistakeChance int // Percentage (0-100)
	UseStrategy   bool
}

// Config returns the configuration for a difficulty.
func (d Difficulty) Config() DifficultyConfig {

	switch d {

	case Easy:
		return DifficultyConfig{
			MistakeChance: 50,
			UseStrategy:   false,
		}

	case Medium:
		return DifficultyConfig{
			MistakeChance: 25,
			UseStrategy:   true,
		}

	case Hard:
		return DifficultyConfig{
			MistakeChance: 10,
			UseStrategy:   true,
		}

	case Impossible:
		return DifficultyConfig{
			MistakeChance: 0,
			UseStrategy:   true,
		}
	}

	return DifficultyConfig{
		MistakeChance: 25,
		UseStrategy:   true,
	}
}

// ChooseCard selects a playable card based on difficulty.
func (g *Game) ChooseCard(player *Player) (int, bool) {

	cfg := player.Difficulty.Config()

	playable := PlayableCards(g, player)

	if len(playable) == 0 {
		return -1, false
	}

	// Easy/Medium occasionally make mistakes.
	if cfg.MistakeChance > 0 &&
		rand.Intn(100) < cfg.MistakeChance {

		return playable[rand.Intn(len(playable))], true
	}

	return BestPlayableCard(g, player)
}