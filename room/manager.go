package room

import (
	"fmt"
	"sync"
)

// Manager manages multiple game rooms.
type Manager struct {
	mu sync.RWMutex

	rooms map[string]*Room
}

// NewManager creates a new instance of Manager.
func NewManager() *Manager {
	return &Manager{
		rooms: make(map[string]*Room),
	}
}

// CreateRoom creates a new room in the manager.
func (m *Manager) CreateRoom(room *Room) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.rooms[room.Code]; exists {
		return fmt.Errorf("room already exists")
	}

	m.rooms[room.Code] = room

	return nil
}

// DeleteRoom removes a room from the manager by its code.
func (m *Manager) DeleteRoom(code string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.rooms, code)
}


// GetRoom retrieves a room by its code.
func (m *Manager) GetRoom(code string) (*Room, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	room, ok := m.rooms[code]

	return room, ok
}


// RoomExists checks if a room exists by its code.
func (m *Manager) RoomExists(code string) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	_, ok := m.rooms[code]

	return ok
}

// Rooms returns the rooms currently managed.
func (m *Manager) Rooms() []*Room {
	m.mu.RLock()
	defer m.mu.RUnlock()

	rooms := make([]*Room, 0, len(m.rooms))

	for _, room := range m.rooms {
		rooms = append(rooms, room)
	}

	return rooms
}

// GenerateUniqueRoomCode generates a unique room code that does not already exist in the manager.
func (m *Manager) GenerateUniqueRoomCode() string {
	for {
		code := GenerateRoomCode()
		if !m.RoomExists(code) {
			return code
		}
	}
}	

// ListRooms returns the rooms currently managed.
func (m *Manager) ListRooms() []*Room {
	m.mu.RLock()
	defer m.mu.RUnlock()

	rooms := make([]*Room, 0, len(m.rooms))

	for _, room := range m.rooms {
		rooms = append(rooms, room)
	}

	return rooms
}