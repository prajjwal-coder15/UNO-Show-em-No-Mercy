package game

// CheckMercy checks whether a player has reached
// the Mercy Rule limit.
//
// If the player has MercyCardLimit (25) or more cards,
// they are eliminated immediately.
func (g *Game) CheckMercy(player *Player) {

	if player == nil {
		return
	}

	if player.Eliminated {
		return
	}

	if player.HandSize() < MercyCardLimit {
		return
	}

	player.Eliminated = true

	// If only one player remains,
	// winner.go will finish the game.
	g.CheckWinner()
}

// IsMercyReached returns true if the player
// has reached the mercy limit.
func (g *Game) IsMercyReached(player *Player) bool {

	if player == nil {
		return false
	}

	return player.HandSize() >= MercyCardLimit
}

// EliminatePlayer removes a player from the game.
func (g *Game) EliminatePlayer(player *Player) {

	if player == nil {
		return
	}

	player.Eliminated = true

	g.CheckWinner()
}

// RevivePlayer is useful for debugging or testing.
func (g *Game) RevivePlayer(player *Player) {

	if player == nil {
		return
	}

	player.Eliminated = false
}