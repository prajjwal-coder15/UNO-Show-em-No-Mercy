package game

// Deck Information
const (
	DeckSize        = 168
	StartingHand    = 10
	MaxPlayers      = 10
	MinPlayers      = 2
	MercyCardLimit  = 28
)

// StandardColors are the normal UNO colors.
var StandardColors = []CardColor{
	Red,
	Blue,
	Green,
	Yellow,
}

// AllColors includes Wild.
var AllColors = []CardColor{
	Red,
	Blue,
	Green,
	Yellow,
	Wild,
}


// NumberValues lists the numeric card values.
var NumberValues = []CardValue{
	One,
	Two,
	Three,
	Four,
	Five,
	Six,
	Seven,
	Eight,
	Nine,
}

// ActionValues lists the classic action card values.
var ActionValues = []CardValue{
	Skip,
	Reverse,
	Draw2,
}

// NoMercyActionValues lists the no mercy action card values.
var NoMercyActionValues = []CardValue{
	DiscardAll,
	SkipEveryone,
}

// WildValues lists the wild card values.
var WildValues = []CardValue{
	Draw4,
	WildDraw6,
	WildDraw10,
	WildReverseDraw4,
	WildColorRoulette,
}

// StackCards lists the card values that can stack and their required draw amounts.
var StackCards = map[CardValue]int{
	Draw2:            2,
	Draw4:            4,
	WildDraw6:        6,
	WildDraw10:       10,
	WildReverseDraw4: 4,
}

// NumberCardSet is a lookup table for number cards.
var NumberCardSet = map[CardValue]struct{}{
	Zero:  {},
	One:   {},
	Two:   {},
	Three: {},
	Four:  {},
	Five:  {},
	Six:   {},
	Seven: {},
	Eight: {},
	Nine:  {},
}

// ActionCardSet is a lookup table for action cards.
var ActionCardSet = map[CardValue]struct{}{
	Skip:         {},
	Reverse:      {},
	Draw2:        {},
	DiscardAll:   {},
	SkipEveryone: {},
}

// WildCardSet is a lookup table for wild cards.
var WildCardSet = map[CardValue]struct{}{
	Draw4:             {},
	WildDraw6:         {},
	WildDraw10:        {},
	WildReverseDraw4:  {},
	WildColorRoulette: {},
}