package game

import "errors"

// SwapHands swaps the hands of two players.
func (g *Game) SwapHands(first, second int) error {

	if first < 0 || first >= len(g.Players) {
		return errors.New("invalid first player index")
	}

	if second < 0 || second >= len(g.Players) {
		return errors.New("invalid second player index")
	}

	if first == second {
		return nil
	}

	g.Players[first].Hand,
		g.Players[second].Hand =
		g.Players[second].Hand,
		g.Players[first].Hand

	return nil
}

// RotateHands rotates every active player's hand.
//
// Clockwise:
// A -> B -> C -> D
//
// becomes
//
// D -> A -> B -> C
func (g *Game) RotateHands() {

	active := make([]*Player, 0)

	for _, player := range g.Players {

		if player.Eliminated {
			continue
		}

		active = append(active, player)
	}

	if len(active) <= 1 {
		return
	}

	lastHand := active[len(active)-1].Hand

	for i := len(active) - 1; i > 0; i-- {
		active[i].Hand = active[i-1].Hand
	}

	active[0].Hand = lastHand
}

// DiscardAllColor removes every card of the specified
// color from the player's hand.
//
// Returns the number of discarded cards.
func (g *Game) DiscardAllColor(
	player *Player,
	color CardColor,
) int {

	if player == nil {
		return 0
	}

	remaining := make([]Card, 0, len(player.Hand))

	discarded := 0

	for _, card := range player.Hand {

		if card.Color == color {

			g.AddToDiscard(card)

			discarded++

			continue
		}

		remaining = append(remaining, card)
	}

	player.Hand = remaining

	return discarded
}

// SkipEveryone skips every remaining player
// and gives the current player another turn.
func (g *Game) SkipEveryone() {

	// Intentionally empty.
	//
	// HumanTurn() / BotTurn() should simply
	// not call AdvanceTurn() after this effect.
}

// LargestHandPlayer returns the index of the active player
// with the largest hand, or -1 if no active players remain.
func (g *Game) LargestHandPlayer() int {

	target := -1
	maxCards := -1

	for i := range g.Players {

		if g.Players[i].Eliminated {
			continue
		}

		if len(g.Players[i].Hand) > maxCards {
			maxCards = len(g.Players[i].Hand)
			target = i
		}
	}

	return target
}