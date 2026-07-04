package game

import "fmt"

// CardColor represents the color of a card.
type CardColor string

//
const (
	Red    CardColor = "Red"
	Blue   CardColor = "Blue"
	Green  CardColor = "Green"
	Yellow CardColor = "Yellow"
	Wild   CardColor = "Wild"
)

// CardValue represents the value or action of a card.
type CardValue string

const (

	// Zero is a number card with value 0.
	Zero CardValue = "0"
	// One is a number card with value 1.
	One CardValue = "1"
	// Two is a number card with value 2.
	Two CardValue = "2"
	// Three is a number card with value 3.
	Three CardValue = "3"
	// Four is a number card with value 4.
	Four CardValue = "4"
	// Five is a number card with value 2.
	Five CardValue = "5"
	// Six is a number card with value 2.
	Six CardValue = "6"
	// Seven is a number card with value 2.
	Seven CardValue = "7"
	// Eight is a number card with value 2.
	Eight CardValue = "8"
	// Nine is a number card with value 2.
	Nine CardValue = "9"

	
	
	// Skip is a skip action card.
	Skip CardValue = "Skip"
	// Reverse is a reverse action card.
	Reverse CardValue = "Reverse"
	// Draw2 is a draw action card that makes the next player draw 2 cards.
	Draw2 CardValue = "Draw2"
	// Draw4 is a draw action card that makes the next player draw 4 cards.
	Draw4 CardValue = "Draw4"

	// -------------------------
	// No Mercy Cards
	// -------------------------

	// SkipEveryone is a card that skips every other player.
	SkipEveryone      CardValue = "SkipEveryone"
	// DiscardAll is a card that discards all cards from the next player's hand.
	DiscardAll        CardValue = "DiscardAll"
	// WildDraw6 is a wild draw action card that makes the next player draw 6 cards.
	WildDraw6         CardValue = "WildDraw6"
	// WildDraw10 is a wild draw action card that makes the next player draw 10 cards.
	WildDraw10        CardValue = "WildDraw10"
	// WildReverseDraw4 is a wild draw action card that makes the next player draw 4 cards and reverses direction.
	WildReverseDraw4  CardValue = "WildReverseDraw4"
	// WildColorRoulette is a wild card that randomly changes to a color (color roulette).
	WildColorRoulette CardValue = "WildColorRoulette"
)

// Card represents one UNO card.
type Card struct {
	Color CardColor
	Value CardValue
}

// String prints a card nicely.
func (c Card) String() string {
	return fmt.Sprintf("%s %s", c.Color, c.Value)
}

// -------------------------
// Helper Methods
// -------------------------

// IsWild returns true if the card is a wild card.
func (c Card) IsWild() bool {
	return c.Color == Wild
}

// IsNumber returns true if the card is a number card.
func (c Card) IsNumber() bool {
	switch c.Value {
	case Zero, One, Two, Three, Four,
		Five, Six, Seven, Eight, Nine:
		return true
	}

	return false
}

// IsDraw returns true if the card causes players to draw cards.
func (c Card) IsDraw() bool {
	switch c.Value {
	case Draw2,
		Draw4,
		WildDraw6,
		WildDraw10,
		WildReverseDraw4:
		return true
	}

	return false
}

// IsAction returns true for non-number colored action cards.
func (c Card) IsAction() bool {
	switch c.Value {
	case Skip,
		Reverse,
		Draw2,
		SkipEveryone,
		DiscardAll:
		return true
	}

	return false
}

// DrawAmount returns how many cards this card adds to the draw stack.
func (c Card) DrawAmount() int {

	switch c.Value {

	case Draw2:
		return 2

	case Draw4:
		return 4

	case WildReverseDraw4:
		return 4

	case WildDraw6:
		return 6

	case WildDraw10:
		return 10
	}

	return 0
}
