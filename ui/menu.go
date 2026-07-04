package ui

import (
	"fmt"
	"strings"
)

// MenuOption represents a single option in the menu.
type MenuOption struct {
	Number int
	Title  string
}

// ShowMainMenu displays the main menu to the user.
func ShowMainMenu() {
	options := []MenuOption{
		{1, "Create Multiplayer Room"},
		{2, "Play With Bots"},
		{3, "Join Room"},
		{4, "Rules"},
		{5, "Exit"},
	}

	PrintMenu(
		"UNO SHOW 'EM NO MERCY",
		options,
	)
}

// PrintMenu prints the menu with the given title and options.
func PrintMenu(
	title string,
	options []MenuOption,
) {
	const width = 34

	border := strings.Repeat("=", width)

	fmt.Println()
	fmt.Println(border)
	fmt.Printf(" %-30s\n", title)
	fmt.Println(border)
	fmt.Println()

	for _, option := range options {
		fmt.Printf("%d. %s\n\n", option.Number, option.Title)
	}

	fmt.Print("> ")
}