package room

import (
	"crypto/rand"
	"math/big"
	"strings"
)

const (
	roomCodeLength = 6

	// Removed confusing characters:
	// O 0 I 1 L
	roomCodeCharset = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"
)

// GenerateRoomCode generates a random room code.
//
// Example:
//
//	X7P9KD
func GenerateRoomCode() string {
	code := make([]byte, roomCodeLength)

	for i := range code {
		n, err := rand.Int(
			rand.Reader,
			big.NewInt(int64(len(roomCodeCharset))),
		)

		if err != nil {
			panic(err)
		}

		code[i] = roomCodeCharset[n.Int64()]
	}

	return string(code)
}

// IsValidRoomCode validates a room code.
func IsValidRoomCode(code string) bool {
	code = strings.TrimSpace(strings.ToUpper(code))

	if len(code) != roomCodeLength {
		return false
	}

	for _, c := range code {
		if !strings.ContainsRune(roomCodeCharset, c) {
			return false
		}
	}

	return true
}

// NormalizeRoomCode converts input to uppercase
// and trims surrounding spaces.
func NormalizeRoomCode(code string) string {
	return strings.ToUpper(
		strings.TrimSpace(code),
	)
}