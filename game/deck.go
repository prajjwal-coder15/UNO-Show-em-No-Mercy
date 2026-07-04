package game

import (
	"errors"
	"math/rand"
	"time"
)

// Deck represents a stack of cards.
type Deck struct {
	Cards []Card
}

// NewDeck creates and returns a fully shuffled
// UNO Show 'Em No Mercy deck.
func NewDeck() *Deck {

	deck := &Deck{
		Cards: make([]Card, 0, 168),
	}

	deck.build()

	deck.Shuffle()

	return deck
}

// build creates the complete deck.
func (d *Deck) build() {

	d.Cards = d.Cards[:0]

	// -----------------------------------
	// Colored Cards
	// -----------------------------------

	for _, color := range StandardColors {

		// One zero
		d.Cards = append(d.Cards, Card{
			Color: color,
			Value: Zero,
		})

		// Two of each number (1-9)
		for _, value := range NumberValues {

			d.Cards = append(d.Cards,
				Card{Color: color, Value: value},
				Card{Color: color, Value: value},
			)
		}

		// Two of each action card
		for _, value := range ActionValues {

			d.Cards = append(d.Cards,
				Card{Color: color, Value: value},
				Card{Color: color, Value: value},
			)
		}

		// No Mercy Colored Cards
		for _, value := range NoMercyActionValues {

			d.Cards = append(d.Cards,
				Card{Color: color, Value: value},
			)
		}
	}

	// -----------------------------------
	// Wild Cards
	// -----------------------------------

	for _, value := range WildValues {

		for i := 0; i < 4; i++ {

			d.Cards = append(d.Cards, Card{
				Color: Wild,
				Value: value,
			})
		}
	}
}

// Shuffle randomizes the deck.
func (d *Deck) Shuffle() {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	r.Shuffle(len(d.Cards), func(i, j int) {

		d.Cards[i], d.Cards[j] =
			d.Cards[j], d.Cards[i]
	})
}

// Draw removes and returns the top card.
func (d *Deck) Draw() (Card, error) {

	if d.Empty() {
		return Card{}, errors.New("deck is empty")
	}

	card := d.Cards[0]

	d.Cards = d.Cards[1:]

	return card, nil
}

// AddBottom inserts cards at the bottom.
func (d *Deck) AddBottom(cards ...Card) {

	d.Cards = append(d.Cards, cards...)
}

// AddTop inserts cards on the top.
func (d *Deck) AddTop(cards ...Card) {

	d.Cards =
		append(cards, d.Cards...)
}

// Peek returns the top card.
func (d *Deck) Peek() (Card, error) {

	if d.Empty() {
		return Card{}, errors.New("deck is empty")
	}

	return d.Cards[0], nil
}

// Empty returns true if no cards remain.
func (d *Deck) Empty() bool {

	return len(d.Cards) == 0
}

// Size returns remaining cards.
func (d *Deck) Size() int {

	return len(d.Cards)
}

// Clear removes every card.
func (d *Deck) Clear() {

	d.Cards = d.Cards[:0]
}