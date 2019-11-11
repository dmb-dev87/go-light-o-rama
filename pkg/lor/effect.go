package lor

// Effect represents a pre-defined action within the LOR hardware.
type Effect byte

const (
	// EffectTwinkle creates random fades with a slight delay.
	EffectTwinkle Effect = 0x06
	// EffectShimmer creates rapid flashing with fading between flashes.
	EffectShimmer Effect = 0x07
)
