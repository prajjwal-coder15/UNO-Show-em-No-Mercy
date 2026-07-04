package ui

import (
	"fmt"

	"uno/room"
)

// PrintRoom prints the room details to the console.
func PrintRoom(room *room.Room) {
	PrintHeader("ROOM")

	fmt.Printf("Code        : %s\n", room.Code)
	fmt.Printf("Name        : %s\n", room.Name)
	fmt.Printf("Host        : %s\n", room.HostID)
	fmt.Printf("Players     : %d/%d\n",
		len(room.Players),
		room.MaxPlayers,
	)

	if room.PlayWithBots {
		fmt.Println("Mode        : Play With Bots")
	} else {
		fmt.Println("Mode        : Multiplayer")
	}

	PrintSeparator()
}

// PrintPlayers prints the list of players in the room to the console.
func PrintPlayers(room *room.Room) {
	PrintSection("Players")

	for i, player := range room.Players {

		status := ""

		if player.Ready {
			status = "✓ Ready"
		} else {
			status = "Waiting"
		}

		host := ""

		if player.ID == room.HostID{
			host = " (Host)"
		}

		bot := ""

		if player.IsBot {
			bot = " 🤖"
		}

		fmt.Printf(
			"%d. %s%s%s - %s\n",
			i+1,
			player.Name,
			host,
			bot,
			status,
		)
	}

	PrintSeparator()
}

// PrintLobby prints the lobby details to the console, including room and player information.
func PrintLobby(room *room.Room) {
	PrintRoom(room)
	PrintPlayers(room)

	fmt.Println()

	if room.PlayWithBots {
		fmt.Println("Press ENTER to start...")
	} else {
		fmt.Println("Waiting for players...")
	}
}