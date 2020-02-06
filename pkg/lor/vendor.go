package lor

import "time"

// DefaultHeartbeatRate represents the timing rate at which LOR vendor software sends a heartbeat to the hardware.
// This has been determined by monitoring the serial port connection.
const DefaultHeartbeatRate = time.Millisecond * 500

var magicOffsetTable = map[byte]byte{
	8:  0x30,
	16: 0x10,
}
