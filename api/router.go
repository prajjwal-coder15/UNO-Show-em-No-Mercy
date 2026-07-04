package api

import "net/http"

// RegisterRoutes registers every HTTP endpoint.
func RegisterRoutes(mux *http.ServeMux) {

	mux.Handle(
		"/health",
		Chain(
			http.HandlerFunc(Health),
			Recover,
			Logging,
			CORS,
			JSON,
		),
	)

	mux.HandleFunc("/bot/move", BotMove)

	mux.HandleFunc("/game/create", CreateGame)
	mux.HandleFunc("/game/start", StartGame)
	mux.HandleFunc("/game/status", GameStatus)

	mux.HandleFunc("/move/play", PlayMove)
	mux.HandleFunc("/move/draw", DrawCard)

	mux.HandleFunc("/rules", Rules)

	// WebSocket (future)
	mux.HandleFunc("/ws", WebSocketHandler)
}