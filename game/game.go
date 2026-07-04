package game

import (
	"errors"
	"uno/rules"
)

// Game represents one UNO Show 'Em No Mercy match.
type Game struct {

	// -------------------------
	// Players
	// -------------------------

	Players []*Player

	CurrentTurn int
	Direction   int

	// -------------------------
	// Cards
	// -------------------------

	DrawPile    *Deck
	DiscardPile []Card

	// -------------------------
	// Current State
	// -------------------------

	ChosenColor CardColor

	PendingDraw int

	LastStackCard Card
	ExtraTurn bool
	SkipCount int
	SkipNextTurn bool

	Started bool
	Finished bool

	Winner *Player

	Rules rules.Config
}

// MoveProvider defines an interface for providing moves in the game.
type MoveProvider interface {
	Move(*Game, *Player) bool
}

// ------------------------------------------------
// Constructor
// ------------------------------------------------

// NewGame creates a new game.
func NewGame(players []*Player) (*Game, error) {

	if len(players) < MinPlayers {
		return nil, errors.New("not enough players")
	}

	if len(players) > MaxPlayers {
		return nil, errors.New("too many players")
	}

	g := &Game{

		Players: players,

		Direction: 1,

		DrawPile: NewDeck(),

		DiscardPile: make([]Card, 0),

		CurrentTurn: 0,
	}

	return g, nil
}

// ------------------------------------------------
// Initialization
// ------------------------------------------------

// Start starts a new game.
func (g *Game) Start() error {

	if g.Started {
		return errors.New("game already started")
	}

	g.DealCards()

	if err := g.StartDiscardPile(); err != nil {
		return err
	}

	g.Started = true

	return nil
}

// ------------------------------------------------
// Current Player
// ------------------------------------------------

// CurrentPlayer returns the player whose turn it is.
func (g *Game) CurrentPlayer() *Player {

	return g.Players[g.CurrentTurn]
}

// ------------------------------------------------
// State
// ------------------------------------------------

// IsFinished returns true if game ended.
func (g *Game) IsFinished() bool {

	return g.Finished
}

// Finish ends the match.
func (g *Game) Finish(player *Player) {

	g.Winner = player

	g.Finished = true
}

// ------------------------------------------------
// Turn Helpers
// ------------------------------------------------

// NextIndex returns next active player.
func (g *Game) NextIndex() int {

	index := g.CurrentTurn

	for {

		index += g.Direction

		if index >= len(g.Players) {
			index = 0
		}

		if index < 0 {
			index = len(g.Players) - 1
		}

		if !g.Players[index].Eliminated {
			return index
		}
	}
}

// AdvanceTurn moves to the next player.
func (g *Game) AdvanceTurn() {

	// Current player plays again.
	if g.ExtraTurn {
		g.ExtraTurn = false
		return
	}

	// Skip players.
	for g.SkipCount > 0 {

		g.CurrentTurn = g.NextIndex()

		g.SkipCount--
	}

	g.CurrentTurn = g.NextIndex()
}

// ------------------------------------------------
// Helpers
// ------------------------------------------------

// ActivePlayers returns remaining players.
func (g *Game) ActivePlayers() int {

	count := 0

	for _, player := range g.Players {

		if !player.Eliminated {
			count++
		}
	}

	return count
}

