package game

import "errors"

// -----------------------------------------------------
// Game Errors
// -----------------------------------------------------

var (
	// ErrGameAlreadyStarted indicates the game has already been started.
	ErrGameAlreadyStarted = errors.New("game already started")
	// ErrGameNotStarted indicates the game has not started.
	ErrGameNotStarted = errors.New("game not started")
	// ErrGameFinished indicates the game has already finished.
	ErrGameFinished = errors.New("game already finished")

	// ErrPlayerNil indicates the player is nil.
	ErrPlayerNil = errors.New("player is nil")
	// ErrInvalidPlayer indicates the player is invalid.
	ErrInvalidPlayer = errors.New("invalid player")
	// ErrPlayerEliminated indicates the player has been eliminated.
	ErrPlayerEliminated = errors.New("player is eliminated")

	// ErrNotPlayersTurn indicates it is not the current player's turn.
	ErrNotPlayersTurn = errors.New("not player's turn")
	// ErrTooFewPlayers indicates there are not enough players.
	ErrTooFewPlayers = errors.New("not enough players")
	// ErrTooManyPlayers indicates there are too many players.
	ErrTooManyPlayers = errors.New("too many players")

	// ErrInvalidCard indicates the card is invalid.
	ErrInvalidCard = errors.New("invalid card")
	// ErrCardNotFound indicates the card was not found.
	ErrCardNotFound = errors.New("card not found")
	// ErrInvalidCardIndex indicates the card index is invalid.
	ErrInvalidCardIndex = errors.New("invalid card index")
	// ErrInvalidPlay indicates the play is invalid.
	ErrInvalidPlay = errors.New("invalid play")
	// ErrNoPlayableCard indicates there is no playable card.
	ErrNoPlayableCard = errors.New("no playable card")

	// ErrDeckEmpty indicates the deck is empty.
	ErrDeckEmpty = errors.New("deck is empty")
	// ErrDiscardEmpty indicates the discard pile is empty.
	ErrDiscardEmpty = errors.New("discard pile is empty")
	// ErrNoCardsToShuffle indicates there are no cards available to reshuffle.
	ErrNoCardsToShuffle = errors.New("no cards available to reshuffle")

	// ErrInvalidColor indicates the color is invalid.
	ErrInvalidColor = errors.New("invalid color")

	// ErrNoPendingDraw indicates there is no pending draw.
	ErrNoPendingDraw = errors.New("no pending draw")

	// ErrNotBot indicates the player is not a bot.
	ErrNotBot = errors.New("player is not a bot")
	// ErrBotTurn indicates the current player is a bot.
	ErrBotTurn = errors.New("current player is a bot")
	// ErrHumanTurn indicates the current player is human.
	ErrHumanTurn = errors.New("current player is human")

	// ErrNoWinner indicates the winner has not been decided.
	ErrNoWinner = errors.New("winner not decided")
)
