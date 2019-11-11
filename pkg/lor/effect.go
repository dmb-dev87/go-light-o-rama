package lor

// Effect represents a pre-defined action within the LOR hardware.
type Effect byte

const (
	// EffectTwinkle creates random fades with a slight delay.
	EffectTwinkle Effect = flagTwinkle
	// EffectShimmer creates rapid flashing with fading between flashes.
	EffectShimmer Effect = flagShimmer
)
