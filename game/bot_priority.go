package game

// CardPriority returns the base priority of a card.
// Higher value = higher priority.
func CardPriority(card Card) int {

	switch card.Value {

	// -------------------------
	// Highest Priority
	// -------------------------

	case WildDraw10:
		return 100

	case WildDraw6:
		return 90

	case WildReverseDraw4:
		return 85

	// -------------------------
	// Strong Specials
	// -------------------------

	case DiscardAll:
		return 80

	case SkipEveryone:
		return 75

	case Draw4:
		return 70

	case Draw2:
		return 60

	// -------------------------
	// Action Cards
	// -------------------------

	case Skip:
		return 50

	case Reverse:
		return 45

	// -------------------------
	// Number Modifiers
	// -------------------------

	case Seven:
		return 35

	case Zero:
		return 30

	// -------------------------
	// Number Cards
	// -------------------------

	case Nine:
		return 9

	case Eight:
		return 8

	case Six:
		return 7

	case Five:
		return 6

	case Four:
		return 5

	case Three:
		return 4

	case Two:
		return 3

	case One:
		return 2

	default:
		return 1
	}
}

// IsHighPriority reports whether a card is considered
// a powerful action card.
func IsHighPriority(card Card) bool {

	switch card.Value {

	case WildDraw10,
		WildDraw6,
		WildReverseDraw4,
		DiscardAll,
		SkipEveryone:

		return true
	}

	return false
}

// PlayableCards returns every playable card index.
func PlayableCards(
	g *Game,
	player *Player,
) []int {

	var result []int

	for i, card := range player.Hand {

		if g.IsValidPlay(card) {
			result = append(result, i)
		}
	}

	return result
}

// HasPlayableCard reports whether the bot has
// at least one legal move.
func HasPlayableCard(
	g *Game,
	player *Player,
) bool {

	_, ok := BestPlayableCard(g, player)

	return ok
}

// ColorCount returns the number of cards of the given color.
func ColorCount(
	player *Player,
	color CardColor,
) int {

	count := 0

	for _, card := range player.Hand {

		if card.Color == color {
			count++
		}
	}

	return count
}

// StrongestColorScore returns the score of the best color.
func StrongestColorScore(
	player *Player,
) int {

	best := 0

	for _, color := range []CardColor{
		Red,
		Blue,
		Green,
		Yellow,
	} {
		score := 0

		for _, card := range player.Hand {

			if card.Color != color {
				continue
			}

			score += CardPriority(card)
		}

		if score > best {
			best = score
		}
	}

	return best
}

// DrawPriority returns the draw strength of a card.
func DrawPriority(card Card) int {
	switch card.Value {
	case Draw2:
		return 2
	case Draw4:
		return 4
	case WildDraw6:
		return 6
	case WildReverseDraw4:
		return 4
	case WildDraw10:
		return 10
	default:
		return 0
	}
}