package game

import (
	"fmt"
)

// Move executes one complete bot turn.
//
// Returns true if the bot played a card.
// Returns false if it had to draw.
func Move(
	g *Game,
	player *Player,
) bool {

	if g == nil || player == nil {
		return false
	}

	// ------------------------------------
	// Pending Draw
	// ------------------------------------

	if g.PendingDraw > 0 {

		index, ok := g.ChooseCard(player)

		// Bot cannot stack.
		if !ok {

			if err := g.AcceptPendingDraw(player); err != nil {
				fmt.Println(err)
			}

			return false
		}

		card := player.Hand[index]

		if !IsStackCard(card) {

			if err := g.AcceptPendingDraw(player); err != nil {
				fmt.Println(err)
			}

			return false
		}

		if err := g.PlayCard(player, index); err != nil {
			fmt.Println(err)
			return false
		}

		chooseWildColor(g, player, card)

		fmt.Printf(
			"%s stacked %s\n",
			player.Name,
			card,
		)

		return true
	}

	// ------------------------------------
	// Normal Turn
	// ------------------------------------

	index, ok := g.ChooseCard(player)

	// No playable card.
	if !ok {

		card, err := g.DrawOne(player)

		if err != nil {
			fmt.Println(err)
			return false
		}

		fmt.Printf(
			"%s drew %s\n",
			player.Name,
			card,
		)

		// Immediately play if possible.
		if g.IsValidPlay(card) {

			last := player.HandSize() - 1

			if err := g.PlayCard(player, last); err != nil {
				fmt.Println(err)
				return false
			}

			chooseWildColor(g, player, card)

			fmt.Printf(
				"%s immediately played %s\n",
				player.Name,
				card,
			)

			return true
		}

		return false
	}

	card := player.Hand[index]

	if err := g.PlayCard(player, index); err != nil {
		fmt.Println(err)
		return false
	}

	chooseWildColor(g, player, card)

	fmt.Printf(
		"%s played %s\n",
		player.Name,
		card,
	)

	return true
}

// chooseWildColor chooses a color after
// playing a wild card.
func chooseWildColor(
	g *Game,
	player *Player,
	card Card,
) {

	switch card.Value {

	case WildDraw6,
		WildDraw10,
		WildReverseDraw4,
		WildColorRoulette:

		color := ChooseColor(player)

		_ = g.SetChosenColor(color)

		fmt.Printf(
			"%s chose %s\n",
			player.Name,
			color,
		)
	}
}

// NewBot creates a new bot player with the given name.
func NewBot(name string) *Player {
	return &Player{
		ID:   GeneratePlayerID(),
		Name:  name,
		IsBot: true,
		Ready: true,
	}
}