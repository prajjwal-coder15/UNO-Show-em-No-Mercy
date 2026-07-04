package game

import (
	"errors"
)

// HumanTurn executes a human player's turn.
//
// cardIndex:
//
//	-1 -> Draw / Accept pending draw
//	>=0 -> Play selected card
func (g *Game) HumanTurn(cardIndex int) error {

	if g.Finished {
		return errors.New("game has finished")
	}

	player := g.CurrentPlayer()

	if player.IsBot {
		return errors.New("current player is a bot")
	}

	// -----------------------------------
	// Draw / Accept Pending Draw
	// -----------------------------------

	if cardIndex == -1 {

		if g.HasPendingDraw() {

			if err := g.AcceptPendingDraw(player); err != nil {
				return err
			}

			g.AdvanceTurn()

			return nil
		}

		return g.DrawUntilPlayable(player)
	}

	// -----------------------------------
	// Play Card
	// -----------------------------------

	if err := g.PlayCard(player, cardIndex); err != nil {
		return err
	}

	g.AdvanceTurn()

	return nil
}

// NextPlayer returns the next active player.
func (g *Game) NextPlayer() *Player {

	index := g.NextIndex()

	return g.Players[index]
}

// PreviousPlayer returns the previous active player.
func (g *Game) PreviousPlayer() *Player {

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
			return g.Players[index]
		}
	}
}

// ReverseDirection reverses play order.
func (g *Game) ReverseDirection() {

	g.Direction *= -1
}

// SkipNextPlayer skips the next player's turn.
func (g *Game) SkipNextPlayer() {

	g.SkipNextTurn = true
}
