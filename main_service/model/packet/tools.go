package packet

import (
	"net"

	"github.com/KouKouChan/YuriCore/main_service/constant"
)

//BuildHeader 建立数据包通用头部
func BuildHeader(seq uint8, id uint8) []byte {
	header := make([]byte, 5)
	header[0] = constant.PacketTypeSignature
	header[1] = seq
	header[2] = 0
	header[3] = 0
	header[4] = id
	return header
}

//GetNextSeq 获取下一次的seq数据包序号
func GetNextSeq(seq *uint8) uint8 {
	if *seq >= constant.MAXSEQUENCE {
		*seq = 0
		return 0
	}
	(*seq)++
	return *seq
}

//SendPacket 发送数据包
func SendPacket(data []byte, client net.Conn) {
	writeLen(&data)
	client.Write(data)
}

//WriteLen 写入数据长度到数据包通用头部
func writeLen(data *[]byte) {
	headerL := uint16(len(*data)) - constant.HeaderLen
	(*data)[2] = uint8(headerL)
	(*data)[3] = uint8(headerL >> 8)
}
