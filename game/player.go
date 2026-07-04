package game

import (
	"errors"
	"fmt"
	"time"
)

// Player represents one player in the game.
type Player struct {
	ID         string
	Name       string
	Hand       []Card
	IsBot      bool
	Eliminated bool
	
	Difficulty Difficulty
	SaidUNO    bool
	Score      int
	Ready      bool
}

// -----------------------------------------------------
// Constructors
// -----------------------------------------------------

// NewPlayer creates a new player.
func NewPlayer(id, name string, isBot bool) *Player {
	player := Player{
		ID:    id,
		Name:  name,
		IsBot: isBot,
		Hand:  make([]Card, 0, StartingHand),
	}

	switch name {
	case "Bot1":
		player.Difficulty = Easy
	case "Bot2":
		player.Difficulty = Medium
	case "Bot3":
		player.Difficulty = Hard
	default:
		player.Difficulty = Impossible
	}

	return &player
}

// -----------------------------------------------------
// Card Management
// -----------------------------------------------------

func (p *Player) isReady() bool {
	return p.Ready
}

// DrawCard adds one card to the player's hand.
func (p *Player) DrawCard(card Card) {
	p.Hand = append(p.Hand, card)
}

// DrawCards adds multiple cards to the player's hand.
func (p *Player) DrawCards(cards ...Card) {
	p.Hand = append(p.Hand, cards...)
}

// RemoveCard removes a card by index.
func (p *Player) RemoveCard(index int) error {

	if index < 0 || index >= len(p.Hand) {
		return errors.New("invalid card index")
	}

	p.Hand = append(
		p.Hand[:index],
		p.Hand[index+1:]...,
	)

	return nil
}

// RemoveSpecificCard removes the first matching card.
func (p *Player) RemoveSpecificCard(card Card) bool {

	for i, c := range p.Hand {

		if c == card {

			_ = p.RemoveCard(i)

			return true
		}
	}

	return false
}

// ClearHand removes all cards.
func (p *Player) ClearHand() {
	p.Hand = p.Hand[:0]
}

// -----------------------------------------------------
// Information
// -----------------------------------------------------

// HandSize returns the number of cards.
func (p *Player) HandSize() int {
	return len(p.Hand)
}

// HasCards returns true if player still has cards.
func (p *Player) HasCards() bool {
	return len(p.Hand) > 0
}

// IsUNO returns true when player has exactly one card.
func (p *Player) IsUNO() bool {
	return len(p.Hand) == 1
}

// IsOut returns true when player has no cards.
func (p *Player) IsOut() bool {
	return len(p.Hand) == 0
}

// -----------------------------------------------------
// Search Helpers
// -----------------------------------------------------

// FindCard returns the index of the card.
func (p *Player) FindCard(card Card) int {

	for i, c := range p.Hand {

		if c == card {
			return i
		}
	}

	return -1
}

// HasColor returns true if player has a card
// of the specified color.
func (p *Player) HasColor(color CardColor) bool {

	for _, card := range p.Hand {

		if card.Color == color {
			return true
		}
	}

	return false
}

// CountColor counts cards of one color.
func (p *Player) CountColor(color CardColor) int {

	count := 0

	for _, card := range p.Hand {

		if card.Color == color {
			count++
		}
	}

	return count
}

// MostCommonColor returns the player's most common color.
func (p *Player) MostCommonColor() CardColor {

	bestColor := Red
	bestCount := -1

	for _, color := range StandardColors {

		count := p.CountColor(color)

		if count > bestCount {

			bestColor = color
			bestCount = count
		}
	}

	return bestColor
}

// CountValue counts cards of one value.
func (p *Player) CountValue(value CardValue) int {

	count := 0

	for _, card := range p.Hand {

		if card.Value == value {
			count++
		}
	}

	return count
}

// -----------------------------------------------------
// UNO
// -----------------------------------------------------

// ResetUNO resets UNO declaration.
func (p *Player) ResetUNO() {
	p.SaidUNO = false
}

// SayUNO marks that the player declared UNO.
func (p *Player) SayUNO() {
	p.SaidUNO = true
}

// -----------------------------------------------------
// Utility
// -----------------------------------------------------

// CopyHand returns a copy of the player's hand.
func (p *Player) CopyHand() []Card {

	hand := make([]Card, len(p.Hand))

	copy(hand, p.Hand)

	return hand
}

// String implements fmt.Stringer.
func (p *Player) String() string {

	return fmt.Sprintf(
		"%s (%d cards)",
		p.Name,
		len(p.Hand),
	)
}
// Eliminate removes the player from the match.
func (p *Player) Eliminate() {
	p.Eliminated = true
}

// Revive puts the player back into the match.
func (p *Player) Revive() {
	p.Eliminated = false
}

// IsEliminated returns true if the player
// has been eliminated.
func (p *Player) IsEliminated() bool {
	return p.Eliminated
}

// GeneratePlayerID returns a unique player identifier.
func GeneratePlayerID() string {
	return fmt.Sprintf("player_%d", time.Now().UnixNano())
}