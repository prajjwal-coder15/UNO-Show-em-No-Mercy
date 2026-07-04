package room

import (
	"time"

	"uno/game"
)

// Room represents a game room where players can join and play Uno.	
type Room struct {
	// Unique room ID
	ID string

	// Join code (e.g. X7P9KD)
	Code string

	// Display name
	Name string

	// Host player ID
	HostID string

	// Players in the room
	Players []*game.Player

	// Current game (nil until started)
	Game *game.Game

	MaxPlayers int

	PlayWithBots bool

	BotCount int

	// Current room state
	State string

	// Creation time
	CreatedAt time.Time
}

// NewRoom creates a new game room with the specified parameters.
func (m *Manager) NewRoom(
	name string,
	hostID string,
	maxPlayers int,
	playWithBots bool,
	botCount int,
) *Room {

	room := &Room{
		ID:            GenerateRoomID(),
		Code:          m.GenerateUniqueRoomCode(),
		Name:          name,
		HostID:        hostID,
		MaxPlayers:    maxPlayers,
		PlayWithBots:  playWithBots,
		BotCount:      botCount,
		State:         "Waiting",
		CreatedAt:     time.Now(),
	}

	m.CreateRoom(room)

	return room
}


// IsFull checks if the room has reached its maximum player capacity.
func (r *Room) IsFull() bool {
	return len(r.Players) >= r.MaxPlayers
}

// IsEmpty checks if the room has no players.
func (r *Room) IsEmpty() bool {
	return len(r.Players) == 0
}

// IsHost checks if the given player ID belongs to the host of the room.
func (r *Room) IsHost(playerID string) bool {
	return r.HostID == playerID
}