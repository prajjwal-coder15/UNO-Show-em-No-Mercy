package rules

import(
	"fmt"
	"strings"
)

// Config contains every configurable rule
// used by the game engine.
type Config struct {

	// -----------------------------
	// Classic Rules
	// -----------------------------

	Stacking bool
	ForceDrawIfNoStack bool

	// -----------------------------
	// Show 'Em No Mercy
	// -----------------------------

	DrawUntilPlayable bool
	MercyRule bool
	DiscardAll bool
	SkipEveryone bool

	// -----------------------------
	// House Rules
	// -----------------------------

	SevenSwap bool
	ZeroRotate bool

	// Allow playing multiple cards
	// with the same value/symbol.
	MultiCardPlay bool

	// Multiple Reverse cards:
	// Two Reverses = Extra Turn.
	DoubleReverseExtraTurn bool

	// Allow any draw card to stack
	// (+2, +4, +6, +10...)
	UniversalDrawStacking bool

	// While a draw penalty exists,
	// any card may be played.
	AllowAnyCardOnPending bool

	// Wild Color Roulette.
	WildColorRoulette bool

	// -----------------------------
	// Gameplay
	// -----------------------------

	MaxPlayers int
	StartingCards int
}

func (c Config) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Stacking: %t\n", c.Stacking)
	fmt.Fprintf(&b, "Force Draw If No Stack: %t\n", c.ForceDrawIfNoStack)

	fmt.Fprintf(&b, "Draw Until Playable: %t\n", c.DrawUntilPlayable)
	fmt.Fprintf(&b, "Mercy Rule: %t\n", c.MercyRule)
	fmt.Fprintf(&b, "Discard All: %t\n", c.DiscardAll)
	fmt.Fprintf(&b, "Skip Everyone: %t\n", c.SkipEveryone)

	fmt.Fprintf(&b, "7 Swap: %t\n", c.SevenSwap)
	fmt.Fprintf(&b, "0 Rotate: %t\n", c.ZeroRotate)
	fmt.Fprintf(&b, "Multi Card Play: %t\n", c.MultiCardPlay)
	fmt.Fprintf(&b, "Double Reverse Extra Turn: %t\n", c.DoubleReverseExtraTurn)
	fmt.Fprintf(&b, "Universal Draw Stacking: %t\n", c.UniversalDrawStacking)
	fmt.Fprintf(&b, "Allow Any Card On Pending: %t\n", c.AllowAnyCardOnPending)
	fmt.Fprintf(&b, "Wild Color Roulette: %t\n", c.WildColorRoulette)

	fmt.Fprintf(&b, "Max Players: %d\n", c.MaxPlayers)
	fmt.Fprintf(&b, "Starting Cards: %d\n", c.StartingCards)

	return b.String()
}