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

var heartbeatPayload = []byte{0x00, 0xFF, 0x81, 0x56, 0x00}

// SendHeartbeat writes a heartbeat payload to the currently open serial port.
func (c Controller) SendHeartbeat() error {
	_, err := c.port.Write(heartbeatPayload)
	return err
}

func (c Controller) On(ch Channel) error {
	_, err := c.port.Write([]byte{
		0x00,
		c.ID,
		flagOn,
		ch.addr(),
		0x00,
	})
	return err
}

func (c Controller) BulkOn(m *mask) error {
	var payload = []byte{
		0x00,
		c.ID,
		m.offset | flagOn,
	}
	payload = append(payload, m.b...)
	payload = append(payload, 0x00)

	_, err := c.port.Write(payload)
	return err
}

// SetBrightness writes a command payload to set the channel's brightness to the specified value.
func (c Controller) SetBrightness(ch Channel, val float64) error {
	_, err := c.port.Write([]byte{
		0x00,
		c.ID,
		flagSet,
		encodeBrightness(val),
		ch.addr(),
		0x00,
	})
	return err
}

func (c Controller) BulkSetBrightness(m *mask, val float64) error {
	var payload = []byte{
		0x00,
		c.ID,
		m.offset | flagSet,
		encodeBrightness(val),
	}
	payload = append(payload, m.b...)
	payload = append(payload, 0x00)

	_, err := c.port.Write(payload)
	return err
}

// SetEffect writes a command payload to set a channel's active effect.
// This will reset the channel's brightness.
func (c Controller) SetEffect(ch Channel, effect Effect) error {
	_, err := c.port.Write([]byte{
		0x00,
		c.ID,
		byte(effect),
		ch.addr(),
		0x00,
	})
	return err
}

func (c Controller) BulkSetEffect(m *mask, effect Effect) error {
	var payload = []byte{
		0x00,
		c.ID,
		m.offset | byte(effect),
	}
	payload = append(payload, m.b...)
	payload = append(payload, 0x00)

	_, err := c.port.Write(payload)
	return err
}

// Fade writes a command payload to fade a channel's brightness from and to the specified values within the specified duration.
func (c Controller) Fade(ch Channel, from, to float64, dur time.Duration) error {
	t, err := encodeDuration(dur)
	if err != nil {
		return err
	}

	_, err = c.port.Write([]byte{
		0x00,
		c.ID,
		flagFade,
		encodeBrightness(from),
		encodeBrightness(to),
		t[0],
		t[1],
		ch.addr(),
		0x00,
	})
	return err
}

func (c Controller) BulkFade(m *mask, from, to float64, dur time.Duration) error {
	t, err := encodeDuration(dur)
	if err != nil {
		return err
	}

	var payload = []byte{
		0x00,
		c.ID,
		m.offset | flagFade,
		encodeBrightness(from),
		encodeBrightness(to),
		t[0],
		t[1],
	}
	payload = append(payload, m.b...)
	payload = append(payload, 0x00)

	_, err = c.port.Write(payload)
	return err
}

// FadeWithEffect writes a command payload to fade a channel's brightness from and to the specified values within the specified duration.
// The effect will be applied alongside the fade effect.
func (c Controller) FadeWithEffect(ch Channel, from, to float64, dur time.Duration, effect Effect) error {
	t, err := encodeDuration(dur)
	if err != nil {
		return err
	}

	_, err = c.port.Write([]byte{
		0x00,
		c.ID,
		byte(effect),
		ch.addr(),
		flagExtendedStatement,
		flagFade,
		encodeBrightness(from),
		encodeBrightness(to),
		t[0],
		t[1],
		0x00,
	})
	return err
}
