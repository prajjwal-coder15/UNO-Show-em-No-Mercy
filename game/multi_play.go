package game

import( 
	"sort"
	"fmt"
)

// CanPlayMultiple returns true if all cards share
// the same value/symbol.
func CanPlayMultiple(cards []Card) bool {

	if len(cards) <= 1 {
		return true
	}

	value := cards[0].Value

	for _, card := range cards[1:] {
		if card.Value != value {
			return false
		}
	}

	return true
}

// CardsFromIndexes returns cards referenced by the indexes.
func CardsFromIndexes(
	player *Player,
	indexes []int,
) ([]Card, error) {

	cards := make([]Card, 0, len(indexes))

	for _, i := range indexes {

		if i < 0 || i >= len(player.Hand) {
			return nil, ErrInvalidCardIndex
		}

		cards = append(cards, player.Hand[i])
	}

	return cards, nil
}

// PlayCards plays multiple cards in one turn.
//
// House Rule:
// All cards must have the same Value.
func (g *Game) PlayCards(
	player *Player,
	indexes []int,
) error {
	if player == nil {
		return ErrPlayerNil
	}

	if len(indexes) == 0 {
		return ErrInvalidCardIndex
	}

	// -------------------------
	// Build selected cards
	// -------------------------

	cards := make([]Card, 0, len(indexes))

	seen := make(map[int]bool)

	for _, index := range indexes {

		if seen[index] {
			return fmt.Errorf("duplicate card selection")
		}

		seen[index] = true

		if index < 0 || index >= len(player.Hand) {
			return ErrInvalidCardIndex
		}

		cards = append(cards, player.Hand[index])
	}

	

	for _, index := range indexes {

		if index < 0 || index >= len(player.Hand) {
			return ErrInvalidCardIndex
		}

		cards = append(cards, player.Hand[index])
	}

	// -------------------------
	// Same symbol/value?
	// -------------------------

	if !CanPlayMultiple(cards) {
		return ErrInvalidPlay
	}

	// -------------------------
	// Every card must be legal.
	// -------------------------

	for _, card := range cards {

		if !g.IsValidPlay(card) {
			return ErrInvalidPlay
		}
	}

	// -------------------------
	// Remove from hand
	// Remove highest indexes first.
	// -------------------------

	sort.Sort(sort.Reverse(sort.IntSlice(indexes)))

	for _, index := range indexes {

		card := player.Hand[index]

		player.RemoveCard(index)

		g.DiscardPile = append(g.DiscardPile, card)
	}

	// -------------------------
	// Apply effects
	// -------------------------

	if err := g.ApplyMultiEffects(player, cards); err != nil {
		return err
	}

	if player.IsUNO() {
	    fmt.Printf("%s says UNO!\n", player.Name)
    }
	g.CheckMercy(player)
	g.CheckWinner()

	return nil
}
