package game

import "fmt"

// ApplyEffect applies the effect of a played card.
func (g *Game) ApplyEffect(player *Player, card Card) error {

	switch card.Value {

	// -------------------------
	// Skip
	// -------------------------

	case Skip:

		fmt.Println("SKIP!")

		g.SkipNextTurn = true

	// -------------------------
	// Reverse
	// -------------------------

	case Reverse:

		fmt.Println("REVERSE!")

		// Official UNO:
		// Reverse acts as Skip in a 2-player game.
		if g.ActivePlayers() == 2 {
			g.SkipNextTurn = true
			return nil
		}

		g.Direction *= -1

	// -------------------------
	// Draw Cards
	// -------------------------

	case Draw2,
		Draw4,
		WildDraw6,
		WildDraw10,
		WildReverseDraw4:

		fmt.Printf("%s!\n", card.Value)

		g.AddToStack(card)

		if card.Value == WildReverseDraw4 {
			g.Direction *= -1
		}

		if player.IsBot {
			g.BotChooseColor(player)
		}

	// -------------------------
	// Discard All
	// -------------------------

	case DiscardAll:

		fmt.Println("DISCARD ALL!")

		discarded := g.DiscardAllColor(
			player,
			card.Color,
		)

		fmt.Printf(
			"%s discarded %d cards.\n",
			player.Name,
			discarded,
		)

	// -------------------------
	// Skip Everyone
	// -------------------------

	case SkipEveryone:

		fmt.Println("SKIP EVERYONE!")

		// Current player immediately gets another turn.
		g.SkipNextTurn = false

	// -------------------------
	// 7 Swap Hands
	// -------------------------

	case Seven:

		fmt.Println("SWAP HANDS!")

		if player.IsBot {

			target := g.LargestHandPlayer()

			if target != -1 &&
				target != g.CurrentTurn {

				_ = g.SwapHands(
					g.CurrentTurn,
					target,
				)
			}
		}

		// Human chooses target in CLI.

	// -------------------------
	// 0 Rotate Hands
	// -------------------------

	case Zero:

		fmt.Println("ROTATE HANDS!")

		g.RotateHands()

	// -------------------------
	// Wild Color Roulette
	// -------------------------

	case WildColorRoulette:

		if player.IsBot {

			color := ChooseColor(player)

			_ = g.SetChosenColor(color)

			return g.WildColorRoulette(
				g.NextPlayer(),
				color,
			)
		}

		// Human chooses the color in CLI.
	}

	return nil
}
