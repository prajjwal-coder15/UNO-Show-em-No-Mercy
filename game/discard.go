package game

import "errors"

// AddToDiscard places a card on top of the discard pile.
func (g *Game) AddToDiscard(card Card) {

	g.DiscardPile = append(g.DiscardPile, card)
}

// TopCard returns the current top discard card.
func (g *Game) TopCard() Card {

	if len(g.DiscardPile) == 0 {
		return Card{}
	}

	return g.DiscardPile[len(g.DiscardPile)-1]
}

// TopCardSafe returns the top discard card safely.
func (g *Game) TopCardSafe() (Card, error) {

	if len(g.DiscardPile) == 0 {
		return Card{}, errors.New("discard pile is empty")
	}

	return g.DiscardPile[len(g.DiscardPile)-1], nil
}

// RemoveTopDiscard removes and returns the top card.
func (g *Game) RemoveTopDiscard() (Card, error) {

	if len(g.DiscardPile) == 0 {
		return Card{}, errors.New("discard pile is empty")
	}

	last := len(g.DiscardPile) - 1

	card := g.DiscardPile[last]

	g.DiscardPile = g.DiscardPile[:last]

	return card, nil
}

// DiscardCount returns the number of cards
// currently in the discard pile.
func (g *Game) DiscardCount() int {

	return len(g.DiscardPile)
}

// ClearDiscard removes every discarded card.
func (g *Game) ClearDiscard() {

	g.DiscardPile = g.DiscardPile[:0]
}

// HasDiscardCards returns true if the discard
// pile contains at least one card.
func (g *Game) HasDiscardCards() bool {

	return len(g.DiscardPile) > 0
}