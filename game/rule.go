package game

// RuleSet is a set of game rules.
type RuleSet struct {
	OfficialRules         bool
}

// Rules contains the active rule settings for the game.
//
// Rules may be modified to change gameplay behavior.
var Rules = RuleSet{
	OfficialRules:         false,
}