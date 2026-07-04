package room

import (
	"fmt"

	"uno/game"
)

// AddPlayer adds a player to the room.
func (r *Room) AddPlayer(player *game.Player) error {
	if r.IsFull() {
		return fmt.Errorf("room is full")
	}

	r.Players = append(r.Players, player)

	return nil
}


// Player Management

// RemovePlayer removes a player from the room by ID.
func (r *Room) RemovePlayer(playerID string) bool {
	for i, player := range r.Players {

		if player.ID != playerID {
			continue
		}

		r.Players = append(
			r.Players[:i],
			r.Players[i+1:]...,
		)

		// Transfer host if necessary.
		if r.HostID == playerID && len(r.Players) > 0 {
			r.HostID = r.Players[0].ID
		}

		return true
	}

	return false
}

// GetPlayer retrieves a player by ID.
func (r *Room) GetPlayer(playerID string) (*game.Player, bool) {
	for _, player := range r.Players {

		if player.ID == playerID {
			return player, true
		}
	}

	return nil, false
}

// SetReady sets the ready status of a player by ID.
func (r *Room) SetReady(
	playerID string,
	ready bool,
) error {

	player, ok := r.GetPlayer(playerID)

	if !ok {
		return fmt.Errorf("player not found")
	}

	player.Ready = ready

	return nil
}

// AddBots adds bots to the room.
func (r *Room) AddBots(count int) error {
	if !r.PlayWithBots {
		return fmt.Errorf("bots disabled")
	}

	for i := 1; i <= count; i++ {

		bot := game.NewBot(
			fmt.Sprintf("Bot %d", i),
		)

		bot.Ready = true

		r.Players = append(r.Players, bot)
	}

	return nil
}

// StartGame starts a game in the room.
func (r *Room) StartGame() error {
	if !r.CanStart() {
		return fmt.Errorf("cannot start game")
	}

	g, err := game.NewGame(r.Players)
		if err != nil {
		return err
		}
		if err:= g.Start(); err!= nil{
			return err
		}

	r.Game = g

	r.State = "Playing"

	return nil
}

// CanStart checks if the game can be started in the room.
func (r *Room) CanStart() bool {
	if len(r.Players) < 2 {
		return false
	}

	for _, p := range r.Players {
		if !p.Ready {
			return false
		}
	}

	return true
}