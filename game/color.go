package game

import (
	"errors"
	"math/rand"
)

// SetChosenColor sets the active color after a wild card.
func (g *Game) SetChosenColor(color CardColor) error {

	if !IsValidColor(color) {
		return errors.New("invalid color")
	}

	g.ChosenColor = color

	return nil
}

// ResetChosenColor resets the chosen color.
func (g *Game) ResetChosenColor() {
	g.ChosenColor = Wild
}

// CurrentColor returns the currently active color.
func (g *Game) CurrentColor() CardColor {
	return g.ChosenColor
}

// ------------------------------------------------------
// Color Validation
// ------------------------------------------------------

// IsValidColor returns true if the color is a standard
// UNO color.
func IsValidColor(color CardColor) bool {

	switch color {

	case Red,
		Blue,
		Green,
		Yellow:
		return true

	default:
		return false
	}
}

// ------------------------------------------------------
// Random Color
// ------------------------------------------------------

// RandomColor returns a random standard color.
func RandomColor() CardColor {

	return StandardColors[rand.Intn(len(StandardColors))]
}

// ------------------------------------------------------
// Bot
// ------------------------------------------------------

// BotChooseColor automatically chooses the best color
// for a bot.
func (g *Game) BotChooseColor(player *Player) CardColor {

	if player == nil {
		return RandomColor()
	}

	color := player.MostCommonColor()

	g.ChosenColor = color

	return color
}

// ------------------------------------------------------
// Human
// ------------------------------------------------------

// HumanChooseColor sets the color selected by
// a human player.
func (g *Game) HumanChooseColor(color CardColor) error {

	return g.SetChosenColor(color)
}