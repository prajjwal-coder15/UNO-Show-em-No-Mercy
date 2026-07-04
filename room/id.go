package room

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateRoomID returns a unique room ID.
func GenerateRoomID() string {
	b := make([]byte, 8)

	if _, err := rand.Read(b); err != nil {
		panic(err)
	}

	return hex.EncodeToString(b)
}