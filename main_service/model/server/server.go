package server

type Server struct {
	ServerIndex  uint8
	ServerName   string
	ServerStatus uint8
	ServerType   uint8
	Channels     []Channel
}

const (
	MAXSERVERROOM = 0xFFFF
)
