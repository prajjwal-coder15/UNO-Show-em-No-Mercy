package ui

import (
	"fmt"
	"strings"
)

const (
	lineWidth = 50
)

// PrintHeader prints a formatted header with a title.
func PrintHeader(title string) {
	border := strings.Repeat("=", lineWidth)

	fmt.Println()
	fmt.Println(border)
	fmt.Printf(" %s\n", title)
	fmt.Println(border)
	fmt.Println()
}

// PrintSection prints a formatted section with a title.
func PrintSection(title string) {
	fmt.Println()
	fmt.Println(title)
	fmt.Println(strings.Repeat("-", lineWidth))
}

// PrintSuccess prints a success message with a checkmark.
func PrintSuccess(message string) {
	fmt.Printf("✓ %s\n", message)
}

// PrintError prints an error message with an X mark.
func PrintError(message string) {
	fmt.Printf("✗ %s\n", message)
}

// PrintInfo prints an informational message with an info symbol.
func PrintInfo(message string) {
	fmt.Printf("ℹ %s\n", message)
}

// PrintSeparator prints a horizontal separator line.
func PrintSeparator() {
	fmt.Println(strings.Repeat("-", lineWidth))
}

// PrintBlankLine prints a blank line for spacing.
func PrintBlankLine() {
	fmt.Println()
}

// WaitForEnter prompts the user to press ENTER to continue.
func WaitForEnter() {
	fmt.Print("Press ENTER to continue...")
	fmt.Scanln()
}

//package ui

// import (
// 	"fmt"
// 	"strings"

// 	"uno/game"
// )

// // -----------------------------------------------------
// // Game State
// // -----------------------------------------------------

// // PrintGameState prints the current board state.
// func PrintGameState(g *game.Game) {

// 	fmt.Println()
// 	fmt.Println(strings.Repeat("=", 60))

// 	top := g.TopCard()

// 	fmt.Printf("Top Card      : %s\n", top)

// 	if g.ChosenColor != "" {
// 		fmt.Printf("Chosen Color  : %s\n", g.ChosenColor)
// 	}

// 	if g.PendingDraw > 0 {
// 		fmt.Printf("Pending Draw  : %d\n", g.PendingDraw)
// 	}

// 	fmt.Printf("Direction     : %s\n", direction(g.Direction))

// 	fmt.Println(strings.Repeat("=", 60))
// 	fmt.Println()
// }

// // -----------------------------------------------------
// // Current Player
// // -----------------------------------------------------

// // Print prints the current player's turn banner.
// func Print(player *game.Player) {

// 	fmt.Printf("\n%s's Turn\n", player.Name)
// 	fmt.Println(strings.Repeat("-", 40))
// }

// // -----------------------------------------------------
// // Messages
// // -----------------------------------------------------

// // PrintUNO announces that a player has called UNO.
// func PrintUNO(player *game.Player) {

// 	fmt.Printf("\n%s shouts UNO!\n", player.Name)
// }

// // PrintWinner outputs the winning player.
// func PrintWinner(player *game.Player) {

// 	fmt.Println()
// 	fmt.Println(strings.Repeat("=", 60))
// 	fmt.Printf("🏆 Winner: %s\n", player.Name)
// 	fmt.Println(strings.Repeat("=", 60))
// }

// // PrintDraw outputs the number of cards drawn by a player.
// func PrintDraw(player *game.Player, amount int) {

// 	fmt.Printf(
// 		"%s draws %d card(s).\n",
// 		player.Name,
// 		amount,
// 	)
// }

// // PrintSkip announces that a player's turn was skipped.
// func PrintSkip(player *game.Player) {

// 	fmt.Printf(
// 		"%s's turn was skipped.\n",
// 		player.Name,
// 	)
// }

// // PrintError prints an error message.
// func PrintError(err error) {

// 	fmt.Println("Error:", err)
// }

// // PrintMessage prints a generic message.
// func PrintMessage(message string) {

// 	fmt.Println(message)
// }

// // -----------------------------------------------------
// // Helpers
// // -----------------------------------------------------

// func direction(value int) string {

// 	if value == 1 {
// 		return "Clockwise"
// 	}

// 	return "Counter-Clockwise"
// }
