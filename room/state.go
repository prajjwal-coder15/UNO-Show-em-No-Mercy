package room

// State represents the current state of a room.
type State string

const (
	// Waiting for players to join.
	Waiting State = "waiting"

	// Starting indicates the host has started the game and setup is in progress.
	Starting State = "starting"

	// Playing indicates a game is currently running.
	Playing State = "playing"

	// Finished indicates the game has ended.
	Finished State = "finished"

	// Closed indicates the room has been closed.
	Closed State = "closed"
)

// IsWaiting returns true if the room is waiting for players.
func (s State) IsWaiting() bool {
	return s == Waiting
}

// IsStarting returns true if the room is starting.
func (s State) IsStarting() bool {
	return s == Starting
}

// IsPlaying returns true if a game is in progress.
func (s State) IsPlaying() bool {
	return s == Playing
}

// IsFinished returns true if the game has ended.
func (s State) IsFinished() bool {
	return s == Finished
}

// IsClosed returns true if the room has been closed.
func (s State) IsClosed() bool {
	return s == Closed
}