package room

import "uno/game"

// FindPlayer returns the player with the given ID.
func (r *Room) FindPlayer(playerID string) (*game.Player, bool) {
	for _, player := range r.Players {
		if player.ID == playerID {
			return player, true
		}
	}

	return nil, false
}

// FindPlayerByName returns the player with the given name.
func (r *Room) FindPlayerByName(name string) (*game.Player, bool) {
	for _, player := range r.Players {
		if player.Name == name {
			return player, true
		}
	}

	return nil, false
}

// HasPlayer returns true if the player exists in the room.
func (r *Room) HasPlayer(playerID string) bool {
	_, ok := r.FindPlayer(playerID)
	return ok
}

// PlayerCount returns the number of players.
func (r *Room) PlayerCount() int {
	return len(r.Players)
}

// HumanPlayers returns only human players.
func (r *Room) HumanPlayers() []*game.Player {
	var players []*game.Player

	for _, player := range r.Players {
		if !player.IsBot {
			players = append(players, player)
		}
	}

	return players
}

// BotPlayers returns only bot players.
func (r *Room) BotPlayers() []*game.Player {
	var bots []*game.Player

	for _, player := range r.Players {
		if player.IsBot {
			bots = append(bots, player)
		}
	}

	return bots
}