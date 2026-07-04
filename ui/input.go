package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

// ReadLine reads a line of input from standard input.
func ReadLine(prompt string) string {
	fmt.Print(prompt)

	text, _ := reader.ReadString('\n')

	return strings.TrimSpace(text)
}

// ReadMenuChoice reads a menu choice between min and max.
func ReadMenuChoice(min, max int) int {
	for {

		fmt.Print("> ")

		text, _ := reader.ReadString('\n')

		text = strings.TrimSpace(text)

		value, err := strconv.Atoi(text)

		if err != nil {
			fmt.Println("Please enter a number.")
			continue
		}

		if value < min || value > max {
			fmt.Printf("Please choose between %d and %d.\n", min, max)
			continue
		}

		return value
	}
}

// ReadBotCount reads the number of bots to play against.
func ReadBotCount() int {
	return ReadMenuChoice(1, 5)
}

// ReadRoomCode reads a room code from standard input.
func ReadRoomCode() string {

	for {

		code := strings.ToUpper(
			ReadLine("Room Code: "),
		)

		if len(code) == 6 {
			return code
		}

		fmt.Println("Room code must contain exactly 6 characters.")
	}
}

// ReadChoices reads a single card selection.
//
// Examples:
//
//	0  -> Draw
//	1
//	5
//
// Returns:
//	-1 = Draw
//	0+ = Card index (0-based)
func ReadChoices() ([]int, error) {
	fmt.Print("Choose card(s) (0 = draw, e.g. 1 or 1 3 5): ")

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	text = strings.TrimSpace(text)
	if text == "" {
		return nil, fmt.Errorf("empty input")
	}

	parts := strings.Fields(text)

	indexes := make([]int, 0, len(parts))

	for _, part := range parts {

		value, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("invalid number: %s", part)
		}

		// Draw
		if value == 0 {

			if len(parts) > 1 {
				return nil, fmt.Errorf("draw cannot be combined with card selection")
			}

			return []int{-1}, nil
		}

		if value < 1 {
			return nil, fmt.Errorf("invalid card number")
		}

		indexes = append(indexes, value-1)
	}

	return indexes, nil
}