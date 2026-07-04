package game

// OpponentNearWin returns true if any opponent has
// two or fewer cards remaining.
func OpponentNearWin(
	g *Game,
	player *Player,
) bool {

	for _, p := range g.Players {

		if p == player {
			continue
		}

		if p.Eliminated {
			continue
		}

		if p.HandSize() <= 2 {
			return true
		}
	}

	return false
}

// ShouldSaveWild returns true if the bot should
// keep a wild card instead of playing it.
func ShouldSaveWild(
	g *Game,
	player *Player,
) bool {

	if g.PendingDraw > 0 {
		return false
	}

	// If someone is about to win,
	// use every powerful card immediately.
	if OpponentNearWin(g, player) {
		return false
	}

	// Keep the wild if we still have many cards.
	return player.HandSize() > 5
}

// ShouldPlayDrawCard returns whether the bot should
// aggressively play a draw card.
func ShouldPlayDrawCard(
	g *Game,
	player *Player,
) bool {

	if g.PendingDraw > 0 {
		return true
	}

	if OpponentNearWin(g, player) {
		return true
	}

	return false
}

// ShouldUseDiscardAll returns true if playing
// Discard All removes enough cards to be worthwhile.
func ShouldUseDiscardAll(
	player *Player,
	card Card,
) bool {

	count := player.CountColor(card.Color)

	return count >= 3
}

// ShouldRotateHands returns true if rotating hands
// is beneficial.
func ShouldRotateHands(
	g *Game,
	player *Player,
) bool {

	largest := 0

	for _, p := range g.Players {

		if p.Eliminated {
			continue
		}

		if p.HandSize() > largest {
			largest = p.HandSize()
		}
	}

	return player.HandSize() < largest
}

// ShouldSwapHands returns true if swapping hands
// would improve the bot's position.
func ShouldSwapHands(
	g *Game,
	player *Player,
) bool {

	target := -1
	maxCards := player.HandSize()

	for i, p := range g.Players {

		if p == player || p.Eliminated {
			continue
		}

		if p.HandSize() > maxCards {
			maxCards = p.HandSize()
			target = i
		}
	}

	return target != -1
}

// SomeoneHasUNO returns true if an opponent
// has exactly one card.
func SomeoneHasUNO(
	g *Game,
	player *Player,
) bool {

	for _, p := range g.Players {

		if p == player {
			continue
		}

		if p.Eliminated {
			continue
		}

		if p.HandSize() == 1 {
			return true
		}
	}

	return false
}