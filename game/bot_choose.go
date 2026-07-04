package game

// BestPlayableCard returns the index of the best playable
// card in the bot's hand.
//
// Returns (-1, false) if no playable card exists.
func BestPlayableCard(
	g *Game,
	player *Player,
) (int, bool) {

	bestIndex := -1
	bestScore := -1

	for i, card := range player.Hand {

		// Ignore cards that cannot currently be played.
		if !g.IsValidPlay(card) {
			continue
		}

		score := CardPriority(card)

		// -------------------------
		// Strategy Bonuses
		// -------------------------

		// Discard All becomes stronger when it
		// removes many cards.
		if card.Value == DiscardAll {

			score += player.CountColor(card.Color) * 15
		}
		
		if SomeoneHasUNO(g, player) {

			if IsHighPriority(card) {
				score += 100
			}
		}

		// Prefer stronger draw cards while stacking.
		if g.PendingDraw > 0 {

			score += DrawPriority(card) * 10
		}

		// If someone is about to win,
		// prioritize attack cards.
		if OpponentNearWin(g, player) {

			if IsHighPriority(card) {
				score += 50
			}
		}

		if score > bestScore {

			bestScore = score
			bestIndex = i
		}
	}

	if bestIndex == -1 {
		return -1, false
	}

	return bestIndex, true
}

