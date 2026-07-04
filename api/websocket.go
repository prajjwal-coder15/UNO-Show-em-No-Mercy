package api

import (
	"log"
	"net/http"
)

// WebSocketHandler is a placeholder for future
// real-time multiplayer support.
//
// TODO:
//   - Upgrade HTTP connection to WebSocket.
//   - Authenticate players.
//   - Join/leave rooms.
//   - Broadcast game events.
//   - Receive player actions.
func WebSocketHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("WebSocket endpoint is not implemented yet.")

	http.Error(
		w,
		"WebSocket support is not implemented",
		http.StatusNotImplemented,
	)
}