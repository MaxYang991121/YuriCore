package controller

import (
	"errors"
	"net"

	"github.com/KouKouChan/YuriCore/main_service/constant"
	"github.com/KouKouChan/YuriCore/main_service/model/packet"
	"github.com/KouKouChan/YuriCore/utils"
	"github.com/KouKouChan/YuriCore/verbose"
)

type PacketController interface {
	GetHeadPacket(data []byte) packet.PacketHeader
	GetDataPacket(header packet.PacketHeader, data []byte) packet.PacketData
	ReadHead(client net.Conn) ([]byte, error)
	ReadData(client net.Conn, len uint16) ([]byte, error)
}

type packetControllerImpl struct {
}

func NewPacketController() PacketController {
	return &packetControllerImpl{}
}

func (p *packetControllerImpl) GetHeadPacket(data []byte) packet.PacketHeader {
	offset := 0
	return packet.PacketHeader{
		data,
		utils.ReadUint8(data, &offset),
		utils.ReadUint16(data, &offset),
	}
}

func (p *packetControllerImpl) GetDataPacket(header packet.PacketHeader, data []byte) packet.PacketData {
	return packet.PacketData{
		data,
		header.Sequence,
		header.Length,
		data[0],
		1,
	}
}

func (p *packetControllerImpl) ReadHead(client net.Conn) ([]byte, error) {
	SeqBuf := make([]byte, 1)
	headlen := constant.HeaderLen - 1
	head, curlen := make([]byte, headlen), 0
	for {
		n, err := client.Read(SeqBuf)
		if err != nil {
			return head, errors.New("read head signature failed !")
		}
		if n >= 1 && SeqBuf[0] == constant.PacketTypeSignature {
			break
		}
		verbose.DebugInfo(2, "Recived a illegal head sig", SeqBuf[0], "from", client.RemoteAddr().String())
	}
	for {
		n, err := client.Read(head[curlen:])

		if err != nil {
			return head, errors.New("read packet head failed !")
		}
		curlen += n
		if curlen >= headlen {
			break
		}
	}
	return head, nil
}

func (p *packetControllerImpl) ReadData(client net.Conn, len uint16) ([]byte, error) {
	data, curlen := make([]byte, len), 0
	for {
		n, err := client.Read(data[curlen:])
		if err != nil {
			return data, errors.New("read packet data failed !")
		}
		curlen += n
		if curlen >= int(len) {
			break
		}
	}
	return data, nil
}
