package game

import "fmt"

// RevealUntilColor reveals cards until a card of the
// chosen color is found.
//
// All revealed cards that do NOT match the chosen color
// are added to the player's hand.
//
// The first matching card is returned and is NOT added
// to the player's hand because it is played immediately.
func (g *Game) RevealUntilColor(
	player *Player,
	color CardColor,
) (Card, error) {

	for {

		card, err := g.DrawFromDeck()
		if err != nil {
			return Card{}, err
		}

		fmt.Printf("%s revealed %v\n", player.Name, card)

		if card.Color == color {
			return card, nil
		}

		player.DrawCard(card)
	}
}

// WildColorRoulette performs the complete Wild Color
// Roulette effect.
//
// Example:
//
//	Player A plays Wild Color Roulette
//	Chooses RED
//
//	Player B reveals:
//
//	Blue 5
//	Green Skip
//	Yellow 2
//	Red Draw2
//
//	Player B keeps:
//	Blue 5
//	Green Skip
//	Yellow 2
//
//	Player B immediately plays:
//	Red Draw2
func (g *Game) WildColorRoulette(
	player *Player,
	color CardColor,
) error {

	card, err := g.RevealUntilColor(player, color)
	if err != nil {
		return err
	}

	fmt.Printf(
		"%s immediately played %v\n",
		player.Name,
		card,
	)

	// Play the revealed card.
	g.AddToDiscard(card)

	// Update current color if the revealed card
	// is not another wild card.
	if !card.IsWild() {
		g.ChosenColor = card.Color
	}

	// Apply the revealed card's effect.
	if err := g.ApplyEffect(player, card); err != nil {
		return err
	}

	// Check mercy and winner in case the revealed
	// card causes draws or eliminations.
	g.CheckMercy(player)
	g.CheckWinner()

	return nil
}