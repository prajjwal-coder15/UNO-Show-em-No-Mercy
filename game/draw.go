package game

import (
	"errors"
	"fmt"
)

// DealCards gives every player the starting hand.
func (g *Game) DealCards() error {

	if g.DrawPile == nil {
		return errors.New("draw pile is nil")
	}

	for i := 0; i < StartingHand; i++ {

		for _, player := range g.Players {

			card, err := g.DrawPile.Draw()
			if err != nil {
				return err
			}

			player.DrawCard(card)
		}
	}

	return nil
}

// DrawOne draws a single card.
func (g *Game) DrawOne(player *Player) (Card, error) {
	card, err := g.DrawFromDeck()
	if err != nil {
		return Card{}, err
	}

	player.DrawCard(card)
	return card, nil
}

// DrawCards draws multiple cards.
func (g *Game) DrawCards(player *Player, amount int) error {

	for i := 0; i < amount; i++ {

		if _, err := g.DrawOne(player); err != nil {
			return err
		}
	}

	return nil
}

// DrawUntilPlayable draws until a playable card is found.
func (g *Game) DrawUntilPlayable(player *Player) error {

	for {

		card, err := g.DrawOne(player)
		if err != nil {
			return err
		}

		fmt.Printf(
			"%s drew %s\n",
			player.Name,
			card,
		)

		if g.IsValidPlay(card) {
			return nil
		}
	}
}

// StartDiscardPile places the first valid card.
func (g *Game) StartDiscardPile() error {

	for {

		card, err := g.DrawPile.Draw()
		if err != nil {
			return err
		}

		// Don't start with Wild cards.
		if card.IsWild() {
			g.DrawPile.AddBottom(card)
			continue
		}

		g.DiscardPile = append(g.DiscardPile, card)

		g.ChosenColor = card.Color

		return nil
	}
}

// Reshuffle recreates the draw pile from discard pile.
func (g *Game) Reshuffle() error {

	if !g.DrawPile.Empty() {
		return nil
	}

	if len(g.DiscardPile) <= 1 {
		return errors.New("no cards left to reshuffle")
	}

	top := g.TopCard()

	cards := append([]Card{}, g.DiscardPile[:len(g.DiscardPile)-1]...)

	g.DiscardPile = []Card{top}

	g.DrawPile.Cards = append(g.DrawPile.Cards, cards...)

	g.DrawPile.Shuffle()

	return nil
}

// DrawFromDeck draws one card without giving it to a player.
func (g *Game) DrawFromDeck() (Card, error) {

	if err := g.Reshuffle(); err != nil {
		return Card{}, err
	}

	return g.DrawPile.Draw()
}