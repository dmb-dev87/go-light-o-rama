package lor

import "time"

func Heartbeat() []byte {
	return []byte{0x00, 0xFF, 0x81, 0x56, 0x00}
}

func On(unit Unit, ch Channel) []byte {
	return ofCommand(commandOn, unit, ch)
}

func MaskedOn(unit Unit, mask *Mask) []byte {
	return ofMaskedCommand(commandOn|mask.offset, unit, mask)
}

func SetBrightness(unit Unit, ch Channel, val float64) []byte {
	return ofCommand(commandSetBrightness, unit, ch, encodeBrightness(val))
}

func MaskedSetBrightness(unit Unit, mask *Mask, val float64) []byte {
	return ofMaskedCommand(commandSetBrightness|mask.offset, unit, mask, encodeBrightness(val))
}

func SetEffect(unit Unit, ch Channel, effect Effect) []byte {
	return ofCommand(byte(effect), unit, ch)
}

func MaskedSetEffect(unit Unit, mask *Mask, effect Effect) []byte {
	return ofMaskedCommand(byte(effect)|mask.offset, unit, mask)
}

func Fade(unit Unit, ch Channel, from, to float64, dur time.Duration) []byte {
	var time = encodeDuration(dur)
	return ofCommand(commandFade, unit, ch, encodeBrightness(from), encodeBrightness(to), time[0], time[1])
}

func MaskedFade(unit Unit, mask *Mask, from, to float64, dur time.Duration) []byte {
	var time = encodeDuration(dur)
	return ofMaskedCommand(commandFade|mask.offset, unit, mask, encodeBrightness(from), encodeBrightness(to), time[0], time[1])
}

func FadeWithEffect(unit Unit, ch Channel, from, to float64, dur time.Duration, effect Effect) []byte {
	var time = encodeDuration(dur)
	return ofCommand(byte(effect), unit, ch, 0x81, commandFade, encodeBrightness(from), encodeBrightness(to), time[0], time[1])
}

func ofCommand(command byte, unit Unit, ch Channel, meta ...byte) (b []byte) {
	b = make([]byte, 3+len(meta)+2)
	b[0] = 0x00
	b[1] = byte(unit)
	b[2] = command
	copy(b[3:], meta)
	b[3+len(meta)] = ch.addr()
	b[4+len(meta)] = 0x00
	return
}

func ofMaskedCommand(command byte, unit Unit, mask *Mask, meta ...byte) (b []byte) {
	b = make([]byte, 3+len(meta)+len(mask.b)+1)
	b[0] = 0x00
	b[1] = byte(unit)
	b[2] = command
	copy(b[3:], meta)
	copy(b[4+len(meta):], mask.b)
	b[len(b)-1] = 0x00
	return
}
