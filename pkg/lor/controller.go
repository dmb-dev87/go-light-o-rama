package lor

import (
	"github.com/tarm/serial"
	"time"
)

// Controller represents a LOR unit.
// ID is the unit's network ID as externally configured.
type Controller struct {
	ID   byte
	port *serial.Port
}

// OpenPort opens a serial port with the given serial.Config object.
// Once opened, OpenPort will send an initial heartbeat (using #SendHeartbeat) to test the connection.
func (c *Controller) OpenPort(conf *serial.Config) error {
	port, err := serial.OpenPort(conf)
	if err != nil {
		return err
	}

	c.port = port
	return c.SendHeartbeat()
}

// SendHeartbeat writes a heartbeat payload to the currently open serial port.
func (c Controller) SendHeartbeat() error {
	_, err := c.port.Write(heartbeatPayload)
	return err
}

// On writes a command payload to set the channel to 100% brightness.
// This saves 1 byte over a SetBrightness call.
func (c Controller) On(ch Channel) error {
	return c.writeSingleCommand(commandOn, ch)
}

// BulkOn writes a multi command payload to set all masked channels to 100% brightness.
// This saves 1 byte over a BulkSetBrightness call.
func (c Controller) BulkOn(m *Mask) error {
	return c.writeMultiCommand(m.offset|commandOn, m)
}

// SetBrightness writes a command payload to set the channel's brightness to the specified value.
func (c Controller) SetBrightness(ch Channel, val float64) error {
	return c.writeSingleCommand(commandSetBrightness, ch, encodeBrightness(val))
}

// BulkSetBrightness writes a multi command payload to set all masked channels brightness to the specified value.
func (c Controller) BulkSetBrightness(m *Mask, val float64) error {
	return c.writeMultiCommand(m.offset|commandSetBrightness, m, encodeBrightness(val))
}

// SetEffect writes a command payload to set a channel's active effect.
// This will reset the channel's brightness.
func (c Controller) SetEffect(ch Channel, effect Effect) error {
	return c.writeSingleCommand(byte(effect), ch)
}

// BulkSetEffect writes a multi command payload to set all masked channels active effect.
func (c Controller) BulkSetEffect(m *Mask, effect Effect) error {
	return c.writeMultiCommand(m.offset|byte(effect), m)
}

// Fade writes a command payload to fade a channel's brightness from and to the specified values within the specified duration.
func (c Controller) Fade(ch Channel, from, to float64, dur time.Duration) error {
	var t = encodeDuration(dur)
	return c.writeSingleCommand(commandFade, ch, encodeBrightness(from), encodeBrightness(to), t[0], t[1])
}

// BulkFade writes a multi command payload to fade all masked channels brightness from and to the specified values within the specified duration.
func (c Controller) BulkFade(m *Mask, from, to float64, dur time.Duration) error {
	var t = encodeDuration(dur)
	return c.writeMultiCommand(m.offset|commandFade, m, encodeBrightness(from), encodeBrightness(to), t[0], t[1])
}

// FadeWithEffect writes a command payload to fade a channel's brightness from and to the specified values within the specified duration.
// The effect will be applied alongside the fade effect.
func (c Controller) FadeWithEffect(ch Channel, from, to float64, dur time.Duration, effect Effect) error {
	var t = encodeDuration(dur)
	return c.writeSingleCommand(byte(effect), ch, 0x81, commandFade, encodeBrightness(from), encodeBrightness(to), t[0], t[1])
}

func (c Controller) writeSingleCommand(id byte, ch Channel, meta ...byte) error {
	var b = []byte{
		0x00,
		c.ID,
		id,
	}
	b = append(b, meta...)
	b = append(b, ch.addr(), 0x00)

	_, err := c.port.Write(b)
	return err
}

func (c Controller) writeMultiCommand(id byte, m *Mask, meta ...byte) error {
	var b = []byte{
		0x00,
		c.ID,
		id,
	}
	b = append(b, meta...)
	b = append(b, m.b...)
	b = append(b, 0x00)

	_, err := c.port.Write(b)
	return err
}
