package packet

import (
	"github.com/KouKouChan/YuriCore/main_service/constant"
	"github.com/KouKouChan/YuriCore/utils"
)

const ()

type InUDPmsg struct {
	Signature uint8
	UserId    uint32
	PortId    uint16
	IpAddress uint32
	Port      uint16
	Seq       uint8

	PacketData         []byte
	Datalen            int
	CurOffset          int //可能32位
	ParsedSuccessfully bool
}

func (p InUDPmsg) IsHeartbeat() bool {
	return p.Datalen == 6
}

func (dest *InUDPmsg) PraseUDPpacket(data []byte, len int) bool {
	dest.CurOffset = 0
	dest.Signature = utils.ReadUint8(data, &dest.CurOffset)
	if dest.Signature != constant.UdpPacketSignature {
		dest.ParsedSuccessfully = false
		return false
	}
	dest.Datalen = len
	dest.PacketData = data
	if dest.IsHeartbeat() {
	} else {
		dest.UserId = utils.ReadUint32(data, &dest.CurOffset)
		dest.PortId = utils.ReadUint16(data, &dest.CurOffset)
		dest.IpAddress = utils.ReadUint32BE(data, &dest.CurOffset)
		dest.Port = utils.ReadUint16(data, &dest.CurOffset)
		dest.Seq = utils.ReadUint8(data, &dest.CurOffset)
	}
	dest.ParsedSuccessfully = true
	return true
}
