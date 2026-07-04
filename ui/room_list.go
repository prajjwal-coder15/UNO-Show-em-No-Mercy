package ui

import (
	"fmt"
	"uno/room"
)

// PrintRoomList displays a formatted list of available rooms.
func PrintRoomList(rooms []*room.Room) {
	PrintHeader("AVAILABLE ROOMS")

	if len(rooms) == 0 {
		fmt.Println("No rooms available.")
		return
	}

	fmt.Printf("%-8s %-20s %-10s %-10s\n",
		"Code", "Room", "Players", "Status")

	PrintSeparator()

	for _, r := range rooms {
		fmt.Printf(
			"%-8s %-20s %d/%d       %-10s\n",
			r.Code,
			r.Name,
			len(r.Players),
			r.MaxPlayers,
			r.State,
		)
	}
}