package lor

const channelStartAddr = 0x80

type Channel byte

func (c Channel) address() byte {
	return channelStartAddr | byte(c)
}
