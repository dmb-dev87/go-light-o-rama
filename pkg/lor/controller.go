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

func (c Controller) On(ch Channel) error {
	return c.writeSingleCommand(commandOn, ch)
}

func (c Controller) BulkOn(m *mask) error {
	return c.writeMultiCommand(m.offset | commandOn, m)
}

// SetBrightness writes a command payload to set the channel's brightness to the specified value.
func (c Controller) SetBrightness(ch Channel, val float64) error {
	return c.writeSingleCommand(commandSetBrightness, ch, encodeBrightness(val))
}

func (c Controller) BulkSetBrightness(m *mask, val float64) error {
	return c.writeMultiCommand(m.offset|commandSetBrightness, m, encodeBrightness(val))
}

// SetEffect writes a command payload to set a channel's active effect.
// This will reset the channel's brightness.
func (c Controller) SetEffect(ch Channel, effect Effect) error {
	return c.writeSingleCommand(byte(effect), ch)
}

func (c Controller) BulkSetEffect(m *mask, effect Effect) error {
	return c.writeMultiCommand(m.offset | byte(effect), m)
}

// Fade writes a command payload to fade a channel's brightness from and to the specified values within the specified duration.
func (c Controller) Fade(ch Channel, from, to float64, dur time.Duration) error {
	t, err := encodeDuration(dur)
	if err != nil {
		return err
	}

	return c.writeSingleCommand(commandFade, ch, encodeBrightness(from), encodeBrightness(to), t[0], t[1])
}

func (c Controller) BulkFade(m *mask, from, to float64, dur time.Duration) error {
	t, err := encodeDuration(dur)
	if err != nil {
		return err
	}

	return c.writeMultiCommand(m.offset | commandFade, m, encodeBrightness(from), encodeBrightness(to), t[0], t[1])
}

// FadeWithEffect writes a command payload to fade a channel's brightness from and to the specified values within the specified duration.
// The effect will be applied alongside the fade effect.
func (c Controller) FadeWithEffect(ch Channel, from, to float64, dur time.Duration, effect Effect) error {
	t, err := encodeDuration(dur)
	if err != nil {
		return err
	}

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

func (c Controller) writeMultiCommand(id byte, m *mask, meta ...byte) error {
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