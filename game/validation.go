package game

// IsValidPlay returns true if the given card can legally be
// played according to the current game state.
func (g *Game) IsValidPlay(card Card) bool {

	// Wild cards are always playable (unless a pending draw exists).
	if card.IsWild() && g.PendingDraw == 0 {
		return true
	}

	// ----------------------------------------
	// Pending Draw Mode
	// ----------------------------------------

	if g.PendingDraw > 0 {

		// Only draw cards may be stacked.
		if !card.IsDraw() {
			return false
		}

		return CanStack(g.LastStackCard, card)
	}

	// ----------------------------------------
	// Normal Mode
	// ----------------------------------------

	top := g.TopCard()

	// Match color.
	if card.Color == g.ChosenColor {
		return true
	}

	// Match value.
	if card.Value == top.Value {
		return true
	}

	// Wild cards.
	if card.IsWild() {
		return true
	}

	return false
}

// IsPlayableCards returns every playable card
// from the player's hand.
func (g *Game) IsPlayableCards(player *Player) []Card {

	playable := make([]Card, 0)

	for _, card := range player.Hand {

		if g.IsValidPlay(card) {
			playable = append(playable, card)
		}
	}

	return playable
}

// HasPlayableCard returns true if the player has
// at least one playable card.
func (g *Game) HasPlayableCard(player *Player) bool {

	for _, card := range player.Hand {

		if g.IsValidPlay(card) {
			return true
		}
	}

	return false
}

// CanPlayIndex safely checks whether the card at the
// specified hand index can be played.
func (g *Game) CanPlayIndex(player *Player, index int) bool {

	if index < 0 || index >= len(player.Hand) {
		return false
	}

	return g.IsValidPlay(player.Hand[index])
}

// IsValidMultiPlay returns true if all cards in the slice can be played
// according to the current game state.
func (g *Game) IsValidMultiPlay(cards []Card) bool {

	if !CanPlayMultiple(cards) {
		return false
	}

	for _, card := range cards {
		if !g.IsValidPlay(card) {
			return false
		}
	}

	return true
}