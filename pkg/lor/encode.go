package lor

import (
	"encoding/binary"
	"math"
	"time"
)

const (
	brightnessMax = 0x01
	brightnessMin = 0xF0
)

func encodeBrightness(f float64) byte {
	switch {
	case f <= 0:
		return brightnessMin
	case f >= 1:
		return brightnessMax
	default:
		return byte(f*(brightnessMax-brightnessMin) + brightnessMin)
	}
}

const (
	timeMax       float64 = 5099
	timeEmptyByte         = 0x80
	durationMax           = time.Second * 25
	durationMin           = time.Millisecond * 100
)

func encodeDuration(dur time.Duration) []byte {
	if dur > durationMax {
		dur = durationMax
	} else if dur < durationMin {
		dur = durationMin
	}

	var b = make([]byte, 2)

	var val = int(math.Round(timeMax / (dur.Seconds() / 0.1)))

	if val <= 0xFF {
		b[0] = timeEmptyByte
		b[1] = byte(val)
	} else {
		binary.BigEndian.PutUint16(b, uint16(val))
	}

	return b
}
