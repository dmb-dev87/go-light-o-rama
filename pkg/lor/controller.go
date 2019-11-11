package lor

import (
	"github.com/tarm/serial"
	"time"
)

type Controller struct {
	Id   byte
	port *serial.Port
}

func (c *Controller) OpenPort(conf *serial.Config) error {
	port, err := serial.OpenPort(conf)
	if err != nil {
		return err
	}

	c.port = port
	return c.SendHeartbeat()
}

var heartbeatPayload = []byte{0x00, 0xFF, 0x81, 0x56, 0x00}

func (c Controller) SendHeartbeat() error {
	_, err := c.port.Write(heartbeatPayload)
	return err
}

func (c Controller) SetBrightness(ch Channel, val float64) error {
	_, err := c.port.Write([]byte{
		0x00,
		c.Id,
		flagSet,
		encodeBrightness(val),
		ch.address(),
		0x00,
	})
	return err
}

func (c Controller) SetEffect(ch Channel, effect Effect) error {
	_, err := c.port.Write([]byte{
		0x00,
		c.Id,
		byte(effect),
		ch.address(),
		0x00,
	})
	return err
}

func (c Controller) Fade(ch Channel, from float64, to float64, dur time.Duration) error {
	t, err := encodeDuration(dur)
	if err != nil {
		return err
	}

	_, err = c.port.Write([]byte{
		0x00,
		c.Id,
		flagFade,
		encodeBrightness(from),
		encodeBrightness(to),
		t[0],
		t[1],
		ch.address(),
		0x00,
	})
	return err
}

func (c Controller) FadeWithEffect(ch Channel, from float64, to float64, dur time.Duration, effect Effect) error {
	t, err := encodeDuration(dur)
	if err != nil {
		return err
	}

	_, err = c.port.Write([]byte{
		0x00,
		c.Id,
		byte(effect),
		ch.address(),
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
