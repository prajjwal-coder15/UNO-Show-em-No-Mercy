package game

// ChooseColor returns the best color for the bot.
func ChooseColor(player *Player) CardColor {

	if player == nil {
		return Red
	}

	counts := map[CardColor]int{
		Red:    0,
		Blue:   0,
		Green:  0,
		Yellow: 0,
	}

	for _, card := range player.Hand {
		// Ignore wild cards.
		if card.Color != Wild {
			counts[card.Color]++
		}
	}

	bestColor := Red
	bestCount := -1

	for _, color := range []CardColor{
		Red,
		Blue,
		Green,
		Yellow,
	} {

		if counts[color] > bestCount {
			bestCount = counts[color]
			bestColor = color
		}
	}

	return bestColor
}