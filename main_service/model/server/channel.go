package server

type Channel struct {
	ChannelIndex  uint8
	ChannelName   string
	ChannelStatus uint8
	ChannelType   uint8

	Rooms []Room
}

const (
	MAXCHANNELROOM = 0xFF
)
