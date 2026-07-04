// package runner

// import (
// 	"fmt"

// 	"uno/game"
// 	"uno/rules"
// 	"uno/ui"
// 	"uno/room"
// )

// // Run starts one complete UNO match.
// func Run() error {

// 	// ---------------------------------
// 	// Welcome
// 	// ---------------------------------

// 	fmt.Println("===================================")
// 	fmt.Println("      UNO Show 'Em No Mercy")
// 	fmt.Println("===================================")

// 	// ---------------------------------
// 	// Game Setup
// 	// ---------------------------------

// 	playerName, err := ui.ReadPlayerName()
// 	if err != nil {
// 		return err
// 	}

// 	botCount, err := ui.ReadBotCount()
// 	if err != nil {
// 		return err
// 	}

// 	preset, err := ui.ReadRulePreset()
// 	if err != nil {
// 		return err
// 	}

// 	ruleConfig := rules.Load(preset)

// 	// ---------------------------------
// 	// Create Players
// 	// ---------------------------------

// 	players := make([]*game.Player, 0, botCount+1)

// 	players = append(players,
// 		game.NewPlayer("1", playerName, false),
// 	)

// 	for i := 1; i <= botCount; i++ {

// 		id := fmt.Sprintf("%d", i+1)
// 		name := fmt.Sprintf("Bot%d", i)

// 		players = append(players,
// 			game.NewPlayer(id, name, true),
// 		)
// 	}

// 	// ---------------------------------
// 	// Create Game
// 	// ---------------------------------

// 	g, err := game.NewGame(players)
// 	if err != nil {
// 		return err
// 	}

// 	g.Rules = ruleConfig

// 	if err := g.Start(); err != nil {
// 		return err
// 	}

// 	fmt.Printf("\nUsing Rules: %s\n\n", preset)

// 	// ---------------------------------
// 	// Main Game Loop
// 	// ---------------------------------

// 	for !g.IsFinished() {

// 		player := g.CurrentPlayer()

// 		ui.PrintGameState(g)
// 		ui.PrintCurrentPlayer(player)

// 		//--------------------------------
// 		// BOT TURN
// 		//--------------------------------

// 		if player.IsBot {

// 			fmt.Printf("\n%s is thinking...\n", player.Name)

// 			if !game.Move(g, player) {
// 				fmt.Println("Bot could not make a move.")
// 			}

// 			g.AdvanceTurn()
// 			continue
// 		}

// 		//--------------------------------
// 		// HUMAN TURN
// 		//--------------------------------

// 		ui.PrintPlayableHand(g, player)

// 		choices, err := ui.ReadChoices()
// 		if err != nil {
// 			ui.PrintError(err)
// 			continue
// 		}

// 		// Draw
// 		if len(choices) == 1 && choices[0] == -1 {

// 			if err := g.HumanTurn(-1); err != nil {
// 				ui.PrintError(err)
// 				continue
// 			}

// 			continue
// 		}

// 		// Single Card
// 		if len(choices) == 1 {

// 			if err := g.HumanTurn(choices[0]); err != nil {
// 				ui.PrintError(err)
// 				continue
// 			}

// 			continue
// 		}

// 		// Multiple Cards
// 		if err := g.PlayCards(player, choices); err != nil {
// 			ui.PrintError(err)
// 			continue
// 		}

// 		g.AdvanceTurn()
// 	}

// 	// ---------------------------------
// 	// Winner
// 	// ---------------------------------

// 	ui.PrintWinner(g.Winner)

// 	return nil
// }

package runner

import (
	"fmt"
    "uno/game"
	"uno/room"
	"uno/rules"
	"uno/ui"
)

// Runner manages the current room and game state.
type Runner struct {
	manager *room.Manager
	room* room.Room
	// Current running game.
	game *game.Game
	rules *rules.Config
}

// New creates a new Runner instance with a room manager.
func New() *Runner {
	return &Runner{
		manager: room.NewManager(),
	}
}

// Run starts the main loop of the UNO game application.
func (r *Runner) Run() {
	for {
		ui.ShowMainMenu()

		switch ui.ReadMenuChoice(1, 5) {

		case 1:
			r.createMultiplayerRoom()

		case 2:
			r.playWithBots()

		case 3:
			r.joinRoom()

		case 4:
			r.showRules()

		case 5:
			ui.PrintInfo("Goodbye!")
			return
		}
	}
}

func (r *Runner) createMultiplayerRoom() {
	name := ui.ReadLine("Player name: ")
	roomName := ui.ReadLine("Room name: ")

	// Create the room.
	room := r.manager.NewRoom(
		roomName,
		r.room.HostID,    // HostID will be set after creating the player.
		6,     // Maximum players.
		false, // Multiplayer mode.
		0,     // No bots.
	)

	// Create the host player.
	host := game.NewPlayer(
		game.GeneratePlayerID(),
		name,
		false,
	)

	host.Ready = true

	// Set the host.
	room.HostID = host.ID

	// Add the host to the room.
	if err := room.AddPlayer(host); err != nil {
		ui.PrintError(err.Error())
		return
	}

	// Show lobby.
	ui.PrintSuccess("Room created successfully!")
	ui.PrintLobby(room)

	// Wait for players.
	ui.PrintInfo("Waiting for players to join...")
	ui.WaitForEnter()
}

func (r *Runner) playWithBots() {

	name := ui.ReadLine("Player name: ")
	roomName := ui.ReadLine("Room name: ")
	botCount := ui.ReadBotCount()

	room := r.manager.NewRoom(
		roomName,
		room.GenerateRoomCode(),     // Host ID (replace with generated player ID)
		6,
		true,
		botCount,
	)

	host := game.NewPlayer(
		game.GeneratePlayerID(),
		name,
		false,
	)

	host.Ready = true

	if err := room.AddPlayer(host); err != nil {
		ui.PrintError(err.Error())
		return
	}

	if err := room.AddBots(botCount); err != nil {
		ui.PrintError(err.Error())
		return
	}

	ui.PrintLobby(room)

	if err := room.StartGame(); err != nil {
		ui.PrintError(err.Error())
		return
	}

	r.game = room.Game

	r.runGame()
}

func (r *Runner) joinRoom() {

	ui.PrintRoomList(r.manager.ListRooms())

	fmt.Println()
	
	name := ui.ReadLine("Player name: ")
	fmt.Println("enter your code:")
	code := ui.ReadRoomCode()

	// Find room.
	room, ok := r.manager.GetRoom(code)
	if !ok {
		ui.PrintError("Room not found.")
		return
	}

	// Room full.
	if room.IsFull() {
		ui.PrintError("Room is full.")
		return
	}

	// Create player.
	player := game.NewPlayer(
		game.GeneratePlayerID(),
		name,
		false,
	)

	// Player joins not ready.
	player.Ready = false

	// Add player.
	if err := room.AddPlayer(player); err != nil {
		ui.PrintError(err.Error())
		return
	}

	ui.PrintSuccess("Joined room successfully!")

	ui.PrintLobby(room)

	// Save current room.
	r.room = room
}

func (r *Runner) showRules() {
	for {
		ui.PrintHeader("RULE PRESETS")

		fmt.Println("1. Official UNO")
		fmt.Println("2. House Rules")
		fmt.Println("3. UNO Show 'Em No Mercy")
		fmt.Println("4. Back")
		fmt.Println()

		switch ui.ReadMenuChoice(1, 4) {

		case 1:
			ui.PrintHeader("OFFICIAL UNO")
			fmt.Println(rules.Load(rules.OfficialPreset))
			ui.WaitForEnter()

		case 2:
			ui.PrintHeader("HOUSE RULES")
			fmt.Println(rules.Load(rules.HousePreset))
			ui.WaitForEnter()

		case 3:
			ui.PrintHeader("UNO SHOW 'EM NO MERCY")
			fmt.Println(rules.Load(rules.NoMercyPreset))
			ui.WaitForEnter()

		case 4:
			return
		}
	}
}


func (r *Runner) runGame() {
	if r.game == nil {
		ui.PrintError("No active game.")
		return
	}

	ui.PrintHeader("GAME STARTED")

	for !r.game.IsFinished() {

		player := r.game.CurrentPlayer()

		ui.PrintSeparator()
		fmt.Printf("Current Turn: %s\n", player.Name)
		ui.PrintSeparator()

		// -------------------------
		// Bot Turn
		// -------------------------
		if player.IsBot {
			r.botTurn(player)

		} else {
			r.humanTurn(player)
		}

		// Advance to next player.
		r.game.AdvanceTurn()
	}

	ui.PrintHeader("GAME OVER")

	if r.game.Winner != nil {
		fmt.Printf("Winner: %s\n", r.game.Winner.Name)
	}
}

func (r *Runner) humanTurn(player *game.Player) {

	fmt.Println("debug: emtered humanturn")
	ui.PrintHeader(player.Name + "'s Turn")

	// Show current game state.
	ui.PrintInfo("Top Card:")
	fmt.Println(r.game.DiscardPile[len(r.game.DiscardPile)-1])

	fmt.Println()
	ui.PrintHand(player)

	for {
		choices, err := ui.ReadChoices()
		if err != nil {
			ui.PrintError(err.Error())
			continue
		}

		// Draw
		if len(choices) == 1 && choices[0] == -1 {

			card, err := r.game.DrawOne(player)
			if err != nil {
				ui.PrintError(err.Error())
				continue
			}

			fmt.Printf("%s drew %s\n", player.Name, card)
			return
		}
	

		// One card
		if len(choices) == 1 {

			card := player.Hand[choices[0]]

			if err := r.game.PlayCard(player, choices[0]); err != nil {
				ui.PrintError(err.Error())
				continue
			}

			if card.IsWild() {
				r.chooseWildColor()
			}

			return
		}

		// Multiple cards
		if err := r.game.PlayCards(player, choices); err != nil {
			ui.PrintError(err.Error())
			continue
		}

		return
	}
}

func (r *Runner) botTurn(player *game.Player) {
	ui.PrintHeader(player.Name + "'s Turn")

	ui.PrintInfo(player.Name + " is thinking...")

	game.Move(r.game, player)

	top := r.game.TopCard(); 
	fmt.Printf("Top Card: %s\n", top)
	
}


func (r *Runner) chooseWildColor() {

	fmt.Println()
	fmt.Println("Choose a color")
	fmt.Println("1. Red")
	fmt.Println("2. Yellow")
	fmt.Println("3. Green")
	fmt.Println("4. Blue")

	choice := ui.ReadMenuChoice(1, 4)

	var color game.CardColor

	switch choice {

	case 1:
		color = game.Red

	case 2:
		color = game.Yellow

	case 3:
		color = game.Green

	case 4:
		color = game.Blue
	}

	if err := r.game.SetChosenColor(color); err != nil {
		ui.PrintError(err.Error())
	}
}


func (r *Runner) printRules(preset rules.Preset){
	r.printRules(rules.OfficialPreset)
	r.printRules(rules.HousePreset)
	r.printRules(rules.NoMercyPreset)
	cfg := rules.Load(preset)

	ui.PrintHeader(preset.String())

	fmt.Println("Rule Preset")
	fmt.Println("--------------------------------")

	fmt.Printf("Stacking: %t\n", cfg.Stacking)
	fmt.Printf("Force Draw If No Stack: %t\n", cfg.ForceDrawIfNoStack)
	fmt.Printf("Draw Until Playable: %t\n", cfg.DrawUntilPlayable)
	fmt.Printf("Mercy Rule: %t\n", cfg.MercyRule)
	fmt.Printf("Discard All: %t\n", cfg.DiscardAll)
	fmt.Printf("Skip Everyone: %t\n", cfg.SkipEveryone)
	fmt.Printf("7 Swap: %t\n", cfg.SevenSwap)
	fmt.Printf("0 Rotate: %t\n", cfg.ZeroRotate)
	fmt.Printf("Multi Card Play: %t\n", cfg.MultiCardPlay)
	fmt.Printf("Double Reverse Extra Turn: %t\n", cfg.DoubleReverseExtraTurn)
	fmt.Printf("Universal Draw Stacking: %t\n", cfg.UniversalDrawStacking)
	fmt.Printf("Allow Any Card On Pending: %t\n", cfg.AllowAnyCardOnPending)
	fmt.Printf("Wild Color Roulette: %t\n", cfg.WildColorRoulette)
	fmt.Printf("Max Players: %d\n", cfg.MaxPlayers)
	fmt.Printf("Starting Cards: %d\n", cfg.StartingCards)

	fmt.Println()

	ui.WaitForEnter()
}