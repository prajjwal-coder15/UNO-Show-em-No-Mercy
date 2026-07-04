package game

// CheckWinner determines whether the game has ended.
//
// A game ends when:
//
//   - Only one active player remains, OR
//   - A player has played all of their cards.
func (g *Game) CheckWinner() {

	if g.Finished {
		return
	}

	// ------------------------------------------------
	// Win by playing all cards
	// ------------------------------------------------

	for _, player := range g.Players {

		if player.Eliminated {
			continue
		}

		if player.IsOut() {

			g.Winner = player
			g.Finished = true

			return
		}
	}

	// ------------------------------------------------
	// Win by Mercy Rule
	// ------------------------------------------------

	active := make([]*Player, 0)

	for _, player := range g.Players {

		if player.Eliminated {
			continue
		}

		active = append(active, player)
	}

	if len(active) == 1 {

		g.Winner = active[0]
		g.Finished = true
	}
}

// WinnerPlayer returns the current winner.
func (g *Game) WinnerPlayer() *Player {

	return g.Winner
}

// HasWinner returns true if the game has a winner.
func (g *Game) HasWinner() bool {

	return g.Winner != nil
}

// ResetWinner clears the winner.
// Useful when starting a new game.
func (g *Game) ResetWinner() {

	g.Winner = nil
	g.Finished = false
}

// RemainingActivePlayers returns every player
// who has not been eliminated.
func (g *Game) RemainingActivePlayers() []*Player {

	active := make([]*Player, 0)

	for _, player := range g.Players {

		if player.Eliminated {
			continue
		}

		active = append(active, player)
	}

	return active
}