package game

// ApplyMultiEffects resolves the combined effects
// of multiple cards played together.
func (g *Game) ApplyMultiEffects(
	player *Player,
	cards []Card,
) error {

	if len(cards) == 0 {
		return nil
	}

	value := cards[0].Value

	switch value {

	case Skip:

		g.SkipCount += len(cards)

	case Draw2:

		g.PendingDraw += len(cards) * 2

	case WildDraw6:

		g.PendingDraw += len(cards) * 6

	case WildDraw10:

		g.PendingDraw += len(cards) * 10

	case WildReverseDraw4:

		g.PendingDraw += len(cards) * 4

		if len(cards)%2 == 1 {
			g.Direction *= -1
		}

		if len(cards) > 1 {
			g.ExtraTurn = true
		}

	case Reverse:

		// Odd reverses reverse direction.
		if len(cards)%2 == 1 {
			g.Direction *= -1
		}

		// Two or more reverses give another turn.
		if len(cards) >= 2 {
			g.ExtraTurn = true
		}

	case DiscardAll:

		for _, card := range cards {
			g.DiscardAllColor(player, card.Color)
		}

	case SkipEveryone:

		g.ExtraTurn = true

	default:
		// Number cards have no combined effect.
	}

	return nil
}