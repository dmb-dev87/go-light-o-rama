package lor

import (
	"bytes"
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

func encodeDuration(dur time.Duration) ([2]byte, error) {
	if dur > durationMax {
		dur = durationMax
	} else if dur < durationMin {
		dur = durationMin
	}

	var b [2]byte

	var val = int(math.Round(timeMax / (dur.Seconds() / 0.1)))

	if val <= 0xFF {
		b[0] = timeEmptyByte
		b[1] = byte(val)
	} else {
		var buf = new(bytes.Buffer)

		err := binary.Write(buf, binary.BigEndian, uint16(val))
		if err != nil {
			return b, err
		}

		copy(b[:], buf.Bytes())
	}

	return b, nil
}
