package api

// CreateGameRequest is sent by the client.
type CreateGameRequest struct {
	PlayerName string `json:"player_name"`
	Bots       int    `json:"bots"`
	RulePreset string `json:"rule_preset"`
}

// PlayMoveRequest is used when a player
// plays one or more cards.
type PlayMoveRequest struct {
	PlayerID string `json:"player_id"`
	Indexes  []int  `json:"indexes"`
}