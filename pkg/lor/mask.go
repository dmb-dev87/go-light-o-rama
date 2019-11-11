package lor

import (
	"errors"
	"math"
)

var errBadLength = errors.New("bad bitLength")

// Mask represents a variable length bit mask.
type Mask struct {
	offset byte
	b      []byte
}

// Set sets the bit value at the specified index to the specified value.
func (m *Mask) Set(ch Channel, val bool) {
	var bitIndex = byte(ch) % 8
	var byteIndex = int(math.Floor(float64(ch) / 8))

	if val {
		m.b[byteIndex] |= 1 << bitIndex
	} else {
		m.b[byteIndex] &= ^(1 << bitIndex)
	}
}

// SetAll sets all bit values within the mask length to the specified value.
func (m *Mask) SetAll(val bool) {
	for i := 0; i < len(m.b); i++ {
		if val {
			m.b[i] = 0xFF
		} else {
			m.b[i] = 0x00
		}
	}
}

// NewMask returns a new mask sized for the specified bit length.
func NewMask(bitLength byte) (*Mask, error) {
	var offset, ok = magicOffsetTable[bitLength]
	if !ok {
		return nil, errBadLength
	}

	return &Mask{
		offset: offset,
		b:      make([]byte, int(math.Ceil(float64(bitLength)/8))),
	}, nil
}
