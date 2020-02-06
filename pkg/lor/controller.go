package lor

import (
	"io"
	"time"
)

// Unit represents a unique, addressable Controller.
type Unit byte

// Controller represents a LOR unit.
// ID is the unit's network ID as externally configured.
type Controller struct {
	io.Writer
	Unit Unit
}

// NewController returns a new Controller instance using the given io.Writer.
func NewController(unit Unit, writer io.Writer) *Controller {
	return &Controller{
		Writer: writer,
		Unit:   unit,
	}
}

// Heartbeat writes a heartbeat payload to the currently open serial port.
func (c Controller) Heartbeat() (n int, err error) {
	return c.Write(Heartbeat())
}

// On writes a command payload to set the channel to 100% brightness.
// This saves 1 byte over a SetBrightness call.
func (c Controller) On(ch Channel) (n int, err error) {
	return c.Write(On(c.Unit, ch))
}

// MaskedOn writes a multi command payload to set all masked channels to 100% brightness.
// This saves 1 byte over a MaskedSetBrightness call.
func (c Controller) MaskedOn(mask *Mask) (n int, err error) {
	return c.Write(MaskedOn(c.Unit, mask))
}

// SetBrightness writes a command payload to set the channel's brightness to the specified value.
func (c Controller) SetBrightness(ch Channel, val float64) (n int, err error) {
	return c.Write(SetBrightness(c.Unit, ch, val))
}

// MaskedSetBrightness writes a multi command payload to set all masked channels brightness to the specified value.
func (c Controller) MaskedSetBrightness(mask *Mask, val float64) (n int, err error) {
	return c.Write(MaskedSetBrightness(c.Unit, mask, val))
}

// SetEffect writes a command payload to set a channel's active effect.
// This will reset the channel's brightness.
func (c Controller) SetEffect(ch Channel, effect Effect) (n int, err error) {
	return c.Write(SetEffect(c.Unit, ch, effect))
}

// MaskedSetEffect writes a multi command payload to set all masked channels active effect.
func (c Controller) MaskedSetEffect(mask *Mask, effect Effect) (n int, err error) {
	return c.Write(MaskedSetEffect(c.Unit, mask, effect))
}

// Fade writes a command payload to fade a channel's brightness from and to the specified values within the specified duration.
func (c Controller) Fade(ch Channel, from, to float64, dur time.Duration) (n int, err error) {
	return c.Write(Fade(c.Unit, ch, from, to, dur))
}

// MaskedFade writes a multi command payload to fade all masked channels brightness from and to the specified values within the specified duration.
func (c Controller) MaskedFade(mask *Mask, from, to float64, dur time.Duration) (n int, err error) {
	return c.Write(MaskedFade(c.Unit, mask, from, to, dur))
}

// FadeWithEffect writes a command payload to fade a channel's brightness from and to the specified values within the specified duration.
// The effect will be applied alongside the fade effect.
func (c Controller) FadeWithEffect(ch Channel, from, to float64, dur time.Duration, effect Effect) (n int, err error) {
	return c.Write(FadeWithEffect(c.Unit, ch, from, to, dur, effect))
}
