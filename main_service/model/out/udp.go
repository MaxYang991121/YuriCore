package out

import (
	"github.com/KouKouChan/YuriCore/main_service/constant"
	"github.com/KouKouChan/YuriCore/main_service/model/packet"
	"github.com/KouKouChan/YuriCore/utils"
)

func BuildUDPHolepunch(index uint16) []byte {
	buf := make([]byte, 3)
	offset := 0
	utils.WriteUint8(&buf, constant.UdpPacketSignature, &offset)
	utils.WriteUint16(&buf, index, &offset)
	return buf[:offset]
}

func UDPBuild(seq uint8, isHost uint8, userid uint32, ip uint32, port uint16) []byte {
	rst := packet.BuildHeader(seq, constant.PacketTypeUdp)
	buf := make([]byte, 12)
	offset := 0
	utils.WriteUint8(&buf, 1, &offset)
	utils.WriteUint8(&buf, isHost, &offset)
	utils.WriteUint32(&buf, userid, &offset)
	utils.WriteUint32BE(&buf, ip, &offset)
	utils.WriteUint16(&buf, port, &offset)
	rst = utils.BytesCombine(rst, buf)
	return rst
}
