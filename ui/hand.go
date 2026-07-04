package ui

import (
	"fmt"

	"uno/game"
)


// PrintHand prints the player's hand to the console.
func PrintHand(player *game.Player) {
	PrintSection(player.Name + "'s Hand")

	if len(player.Hand) == 0 {
		fmt.Println("No cards.")
		return
	}

	for i, card := range player.Hand {

		fmt.Printf(
			"%2d. %s\n",
			i+1,
			card.String(),
		)
	}

	PrintSeparator()
}







































































// package ui

// import (
// 	"fmt"

// 	"uno/game"
// )

// // PrintPlayableHand prints the player's hand and marks
// // playable cards with a ✓.
// func PrintPlayableHand(
// 	g *game.Game,
// 	player *game.Player,
// ) {

// 	if player == nil {
// 		return
// 	}

// 	fmt.Println()
// 	fmt.Printf("%s's Hand (%d cards)\n", player.Name, player.HandSize())
// 	fmt.Println("------------------------------------------------------------")

// 	for i, card := range player.Hand {

// 		marker := " "

// 		if g.IsValidPlay(card) {
// 			marker = "✓"
// 		}

// 		fmt.Printf(
// 			"%2d. [%s] %s\n",
// 			i+1,
// 			marker,
// 			card,
// 		)
// 	}

// 	fmt.Println("------------------------------------------------------------")
// 	fmt.Println("✓ = Playable")
// 	fmt.Println()
// }