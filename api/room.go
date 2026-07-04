package api

import (
	"fmt"
	"sync"

	"uno/game"
)

// Room represents one active UNO game.
type Room struct {
	ID   string
	Name string
	Code string
	Game *game.Game
	PlayWithBots bool
	BotCount     int
	Status	   string
	MaxPlayers   int
	Players      []*game.Player
}

// RoomManager stores all active rooms.
type RoomManager struct {
	mu    sync.RWMutex
	rooms map[string]*Room
}

// Rooms is the global room manager.
var Rooms = NewRoomManager()

// NewRoomManager creates a new room manager.
func NewRoomManager() *RoomManager {
	return &RoomManager{
		rooms: make(map[string]*Room),
	}
}

// Create adds a new room.
func (rm *RoomManager) Create(g *game.Game) *Room {

	rm.mu.Lock()
	defer rm.mu.Unlock()

	id := fmt.Sprintf("room-%d", len(rm.rooms)+1)

	room := &Room{
		ID:   id,
		Game: g,
	}

	rm.rooms[id] = room

	return room
}

// Get returns a room by ID.
func (rm *RoomManager) Get(id string) (*Room, bool) {

	rm.mu.RLock()
	defer rm.mu.RUnlock()

	room, ok := rm.rooms[id]

	return room, ok
}

// Delete removes a room.
func (rm *RoomManager) Delete(id string) {

	rm.mu.Lock()
	defer rm.mu.Unlock()

	delete(rm.rooms, id)
}

// List returns every active room.
func (rm *RoomManager) List() []*Room {

	rm.mu.RLock()
	defer rm.mu.RUnlock()

	list := make([]*Room, 0, len(rm.rooms))

	for _, room := range rm.rooms {
		list = append(list, room)
	}

	return list
}

// Count returns the number of active rooms.
func (rm *RoomManager) Count() int {

	rm.mu.RLock()
	defer rm.mu.RUnlock()

	return len(rm.rooms)
}

// RoomExists checks whether a room with the given code exists.
func (rm *RoomManager) RoomExists(code string) bool {
	rm.mu.RLock()
	defer rm.mu.RUnlock()
	_, ok := rm.rooms[code]
	return ok
}

// GenerateUniqueRoomCode generates a unique room code that does not exist yet.
func (rm *RoomManager) GenerateUniqueRoomCode() string {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	for {
		code := game.GenerateRoomCode()

		if !rm.RoomExists(code) {
			return code
		}
	}
}