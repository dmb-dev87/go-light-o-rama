package lor

const channelStartAddr = 0x80

type Channel byte

func (c Channel) addr() byte {
	return channelStartAddr | byte(c)
}
