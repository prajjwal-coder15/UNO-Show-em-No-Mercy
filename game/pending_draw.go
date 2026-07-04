package game

import (
	"errors"
	"fmt"
)

// AcceptPendingDraw makes the player accept the current
// draw penalty and resets the draw stack.
func (g *Game) AcceptPendingDraw(player *Player) error {

	if player == nil {
		return errors.New("player is nil")
	}

	if g.PendingDraw == 0 {
		return nil
	}

	amount := g.PendingDraw

	fmt.Printf(
		"%s accepted the draw penalty (%d cards).\n",
		player.Name,
		amount,
	)

	// Draw every pending card.
	if err := g.DrawCards(player, amount); err != nil {
		return err
	}

	// Reset draw stack.
	g.ResetStack()

	// Apply Mercy Rule after drawing.
	g.CheckMercy(player)

	return nil
}

// AddPendingDraw increases the current draw penalty.
func (g *Game) AddPendingDraw(amount int) {

	if amount <= 0 {
		return
	}

	g.PendingDraw += amount
}

// ClearPendingDraw removes the current draw penalty.
func (g *Game) ClearPendingDraw() {

	g.PendingDraw = 0
	g.LastStackCard = Card{}
}

// PendingDrawAmount returns the current draw penalty.
func (g *Game) PendingDrawAmount() int {

	return g.PendingDraw
}

// HasPendingPenalty returns true if a draw penalty exists.
func (g *Game) HasPendingPenalty() bool {

	return g.PendingDraw > 0
}