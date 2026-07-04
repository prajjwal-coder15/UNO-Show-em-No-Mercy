package rules

// House returns your custom house rules.
func House() Config {
	return Config{

		// -----------------------------
		// Classic Rules
		// -----------------------------

		// Allow stacking draw cards.
		Stacking: true,

		// If a player cannot continue the
		// stack, they must take the penalty.
		ForceDrawIfNoStack: true,

		// -----------------------------
		// Show 'Em No Mercy
		// -----------------------------

		DrawUntilPlayable: true,
		MercyRule: true,
		DiscardAll: true,
		SkipEveryone: true,
		WildColorRoulette: true,

		// -----------------------------
		// House Rules
		// -----------------------------

		// Play multiple cards with the
		// same value/symbol.
		MultiCardPlay: true,

		// 7 = Swap hands.
		SevenSwap: true,

		// 0 = Rotate hands.
		ZeroRotate: true,

		// Two or more Reverse cards
		// give another turn.
		DoubleReverseExtraTurn: true,

		// Any draw card may stack with
		// any other draw card.
		//
		// Examples:
		// +2 → +4
		// +4 → +10
		// +6 → +2
		UniversalDrawStacking: true,

		// While a draw penalty exists,
		// only another draw card may be
		// played.
		AllowAnyCardOnPending: false,

		// -----------------------------
		// Game Settings
		// -----------------------------

		MaxPlayers:    10,
		StartingCards: 7,
	}
}