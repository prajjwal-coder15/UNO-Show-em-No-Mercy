package game

import (
	"crypto/rand"
	"math/big"
)

const (
	roomCodeLength = 6

	// Removed confusing characters:
	// 0 O 1 I L
	roomCodeCharset = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"
)

// GenerateRoomCode returns a random room code like:
// X7P9KD
func GenerateRoomCode() string {
	code := make([]byte, roomCodeLength)

	for i := range code {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(roomCodeCharset))))
		if err != nil {
			panic(err)
		}

		code[i] = roomCodeCharset[n.Int64()]
	}

	return string(code)
}

// IsValidRoomCode checks the format of a room code.
func IsValidRoomCode(code string) bool {
	if len(code) != roomCodeLength {
		return false
	}

	for _, c := range code {
		found := false

		for _, allowed := range roomCodeCharset {
			if c == allowed {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	return true
}

// GenerateRoomID returns a random room ID.
func GenerateRoomID() string {
	id, err := rand.Int(rand.Reader, big.NewInt(1<<62))
	if err != nil {
		panic(err)
	}
	return id.String()
}