package packet

type (
	//PacketHeader ,header of packet , 4 bytes len
	PacketHeader struct {
		Data     []byte
		Sequence uint8
		Length   uint16
	}
	//PacketData ,data part of packet
	PacketData struct {
		Data      []byte
		Sequence  uint8
		Length    uint16
		Id        uint8
		CurOffset int
	}
)

