package lor

const channelStartAddr = 0x80

// Channel represents a unique, addressable channel belonging to a given Controller's namespace.
type Channel byte

func (c Channel) addr() byte {
	return channelStartAddr | byte(c)
}
