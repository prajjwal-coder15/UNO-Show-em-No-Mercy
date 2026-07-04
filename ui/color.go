package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"uno/game"
)

// PrintColors displays the available colors.
func PrintColors() {

	fmt.Println()
	fmt.Println("Choose a color:")
	fmt.Println("1. Red")
	fmt.Println("2. Blue")
	fmt.Println("3. Green")
	fmt.Println("4. Yellow")
	fmt.Println()
}

// ReadColor reads a color from the terminal.
func ReadColor() (game.CardColor, error) {

	reader := bufio.NewReader(os.Stdin)

	for {

		PrintColors()

		fmt.Print("Choice: ")

		text, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}

		switch strings.TrimSpace(strings.ToLower(text)) {

		case "1", "r", "red":
			return game.Red, nil

		case "2", "b", "blue":
			return game.Blue, nil

		case "3", "g", "green":
			return game.Green, nil

		case "4", "y", "yellow":
			return game.Yellow, nil

		default:
			fmt.Println("Invalid color. Try again.")
		}
	}
}

// PrintChosenColor displays the selected color.
func PrintChosenColor(color game.CardColor) {

	fmt.Printf(
		"\nChosen Color: %s\n\n",
		color,
	)
}

// PrintBotColorChoice prints the bot's chosen color.
func PrintBotColorChoice(
	playerName string,
	color game.CardColor,
) {

	fmt.Printf(
		"%s chose %s.\n",
		playerName,
		color,
	)
}