package api

import (
	"encoding/json"
	"net/http"

	"uno/game"
)

// BotMoveRequest requests a bot to play.
type BotMoveRequest struct {
	RoomID string `json:"room_id"`
	Player string `json:"player"`
}

// BotMoveResponse is returned after the bot moves.
type BotMoveResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// BotMove executes one bot turn.
func BotMove(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req BotMoveRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	room, ok := Rooms.Get(req.RoomID)
	if !ok {
		http.Error(w, "room not found", http.StatusNotFound)
		return
	}

	g := room.Game

	player := g.CurrentPlayer()

	if player == nil {
		http.Error(w, "current player not found", http.StatusInternalServerError)
		return
	}

	if !player.IsBot {
		http.Error(w, "current player is not a bot", http.StatusBadRequest)
		return
	}

	if !game.Move(g, player) {
		http.Error(w, "bot could not make a move", http.StatusInternalServerError)
		return
	}

	g.AdvanceTurn()

	response := BotMoveResponse{
		Success: true,
		Message: "bot turn completed",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}