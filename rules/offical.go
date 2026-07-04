package rules

// Official returns the official UNO rules.
func Official() Config {
	return Config{

		// -----------------------------
		// Classic UNO
		// -----------------------------

		Stacking:             false,
		ForceDrawIfNoStack:   true,

		// -----------------------------
		// Show 'Em No Mercy
		// -----------------------------

		DrawUntilPlayable:    false,
		MercyRule:            false,
		DiscardAll:           false,
		SkipEveryone:         false,

		// -----------------------------
		// House Rules
		// -----------------------------

		SevenSwap:            false,
		ZeroRotate:           false,
		MultiCardPlay:        false,
		DoubleReverseExtraTurn: false,
		UniversalDrawStacking: false,
		AllowAnyCardOnPending: false,
		WildColorRoulette:    false,

		// -----------------------------
		// Game Settings
		// -----------------------------

		MaxPlayers:    10,
		StartingCards: 7,
	}
}