package game

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// -----------------------------------------------------
// Card Helpers
// -----------------------------------------------------

// IsNumber returns true if the card value is a number.
func IsNumber(value CardValue) bool {
	switch value {
	case Zero,
		One,
		Two,
		Three,
		Four,
		Five,
		Six,
		Seven,
		Eight,
		Nine:
		return true
	}

	return false
}

// IsWild returns true if the card is a wild card.
func IsWild(card Card) bool {
	return card.Color == Wild
}

// IsStackCard returns true if the card participates
// in draw stacking.
func IsStackCard(card Card) bool {
	return card.IsDraw()
}

// -----------------------------------------------------
// Slice Helpers
// -----------------------------------------------------

// ShuffleCards shuffles a slice of cards.
func ShuffleCards(cards []Card) {

	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
}

// CopyCards returns a copy of a card slice.
func CopyCards(cards []Card) []Card {

	result := make([]Card, len(cards))

	copy(result, cards)

	return result
}

// -----------------------------------------------------
// Player Helpers
// -----------------------------------------------------

// ActivePlayersList returns all active players.
func (g *Game) ActivePlayersList() []*Player {

	players := make([]*Player, 0)

	for _, player := range g.Players {

		if player.Eliminated {
			continue
		}

		players = append(players, player)
	}

	return players
}

// NextPlayerIndex returns the next active player's index.
func (g *Game) NextPlayerIndex() int {

	index := g.CurrentTurn

	for {

		index += g.Direction

		if index >= len(g.Players) {
			index = 0
		}

		if index < 0 {
			index = len(g.Players) - 1
		}

		if !g.Players[index].Eliminated {
			return index
		}
	}
}

// PreviousPlayerIndex returns the previous active player's index.
func (g *Game) PreviousPlayerIndex() int {

	index := g.CurrentTurn

	for {

		index -= g.Direction

		if index < 0 {
			index = len(g.Players) - 1
		}

		if index >= len(g.Players) {
			index = 0
		}

		if !g.Players[index].Eliminated {
			return index
		}
	}
}

// -----------------------------------------------------
// Game Helpers
// -----------------------------------------------------

// IsRunning returns true if the game has started
// and not yet finished.
func (g *Game) IsRunning() bool {

	return g.Started && !g.Finished
}

// RemainingPlayers returns the number of players
// still in the game.
func (g *Game) RemainingPlayers() int {

	count := 0

	for _, player := range g.Players {

		if !player.Eliminated {
			count++
		}
	}

	return count
}

// RandomPlayer returns a random active player.
// Returns nil if no active players exist.
func (g *Game) RandomPlayer() *Player {

	active := g.ActivePlayersList()

	if len(active) == 0 {
		return nil
	}

	return active[rand.Intn(len(active))]
}