package game

import (
	"errors"
)

// PlayCard plays the card at the given hand index.
func (g *Game) PlayCard(player *Player, index int) error {

	// -------------------------
	// Validation
	// -------------------------

	if player == nil {
		return errors.New("player is nil")
	}

	if g.Finished {
		return errors.New("game already finished")
	}

	if index < 0 || index >= player.HandSize() {
		return errors.New("invalid card index")
	}

	card := player.Hand[index]

	if !g.IsValidPlay(card) {
		return errors.New("invalid play")
	}

	// -------------------------
	// Remove from hand
	// -------------------------

	if err := player.RemoveCard(index); err != nil {
		return err
	}

	// -------------------------
	// Add to discard pile
	// -------------------------

	g.AddToDiscard(card)

	// -------------------------
	// Update chosen color
	// -------------------------

	if !card.IsWild() {
		g.ChosenColor = card.Color
	}

	// -------------------------
	// Apply card effect
	// -------------------------

	g.ApplyEffect(player, card)

	// -------------------------
	// Mercy Rule
	// -------------------------

	g.CheckMercy(player)

	// -------------------------
	// Winner
	// -------------------------

	if player.IsOut() {
		g.Finish(player)
		return nil
	}

	return nil
}

// PlayCardByValue plays the first matching card.
func (g *Game) PlayCardByValue(player *Player, card Card) error {

	index := player.FindCard(card)

	if index == -1 {
		return errors.New("card not found")
	}

	return g.PlayCard(player, index)
}

// CanPlayCard returns whether a card may be played.
func (g *Game) CanPlayCard(player *Player, index int) bool {

	if index < 0 || index >= player.HandSize() {
		return false
	}

	return g.IsValidPlay(player.Hand[index])
}