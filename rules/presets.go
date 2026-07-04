package rules

// Preset represents a rules preset for UNO.
type Preset string

const (
	// OfficialPreset represents the official UNO rules.
	OfficialPreset Preset = "official"
	// HousePreset represents house rules variant.
	HousePreset Preset = "house"
	// NoMercyPreset represents the no mercy variant.
	NoMercyPreset Preset = "no_mercy"
)

func (p Preset) String() string {
	switch p {
	case OfficialPreset:
		return "Official UNO"

	case HousePreset:
		return "House Rules"

	case NoMercyPreset:
		return "UNO Show 'Em No Mercy"

	default:
		return "Unknown"
	}
}

// Load returns the configuration for the given preset.
func Load(preset Preset) Config {
	switch preset {
	case OfficialPreset:
		return Official()

	case HousePreset:
		return House()

	case NoMercyPreset:
		return NoMercy()

	default:
		return Official()
	}
}

// Available returns the list of supported UNO rule presets.
func Available() []Preset {
	return []Preset{
		OfficialPreset,
		HousePreset,
		NoMercyPreset,
	}
}