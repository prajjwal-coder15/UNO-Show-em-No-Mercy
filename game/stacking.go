package game

// IsStackingActive returns true if a draw stack is currently active.
func (g *Game) IsStackingActive() bool {
	return g.PendingDraw > 0
}

// CanStack returns true if the current draw card
// can be stacked on the previous draw card.
//
// House Rule:
// Any draw card may be stacked on any other draw card.
func CanStack(previous Card, current Card) bool {

	if !previous.IsDraw() {
		return false
	}

	if !current.IsDraw() {
		return false
	}

	return true
}

// AddToStack adds the card's draw amount
// to the current pending draw.
func (g *Game) AddToStack(card Card) {

	if !card.IsDraw() {
		return
	}

	g.PendingDraw += card.DrawAmount()
	g.LastStackCard = card
}

// ResetStack clears the current draw stack.
func (g *Game) ResetStack() {

	g.PendingDraw = 0
	g.LastStackCard = Card{}
}

// StackAmount returns the current pending draw.
func (g *Game) StackAmount() int {
	return g.PendingDraw
}

// HasPendingDraw returns true if there is
// an active draw penalty.
func (g *Game) HasPendingDraw() bool {
	return g.PendingDraw > 0
}