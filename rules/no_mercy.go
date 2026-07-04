package rules

// NoMercy returns the official UNO Show 'Em No Mercy rules.
func NoMercy() Config {
	return Config{

		// -----------------------------
		// Classic Rules
		// -----------------------------

		// Draw cards may be stacked.
		Stacking: true,

		// If a player cannot continue the stack,
		// they immediately take the penalty.
		ForceDrawIfNoStack: true,

		// -----------------------------
		// Show 'Em No Mercy
		// -----------------------------

		// Continue drawing until a playable
		// card is found.
		DrawUntilPlayable: true,

		// Eliminate players with too many cards.
		MercyRule: true,

		// Discard every card of one color.
		DiscardAll: true,

		// Skip every other player.
		SkipEveryone: true,

		// Wild Color Roulette exists in
		// Show 'Em No Mercy.
		WildColorRoulette: true,

		// -----------------------------
		// House Rules
		// -----------------------------

		// These are NOT official rules.
		SevenSwap: false,
		ZeroRotate: false,
		MultiCardPlay: false,
		DoubleReverseExtraTurn: false,

		// Only compatible draw cards may stack.
		UniversalDrawStacking: false,

		// During a pending draw, players cannot
		// play normal cards.
		AllowAnyCardOnPending: false,

		// -----------------------------
		// Game Settings
		// -----------------------------

		MaxPlayers: 10,
		StartingCards: 7,
	}
}
