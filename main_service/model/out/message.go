package out

import (
	"net"

	"github.com/KouKouChan/YuriCore/main_service/constant"
	"github.com/KouKouChan/YuriCore/main_service/model/packet"
	"github.com/KouKouChan/YuriCore/utils"
)

func OnSendMessage(seq uint8, client net.Conn, tp uint8, msg string) {
	buf := make([]byte, 128)
	offset := 0

	ansi, _ := utils.Utf8ToLocal(msg)
	utils.WriteUint8(&buf, tp, &offset)
	utils.WriteStringWithNull(&buf, []byte(ansi), &offset)
	utils.WriteUint8(&buf, 0, &offset)

	packet.SendPacket(utils.BytesCombine(packet.BuildHeader(seq, constant.PacketTypeChat), buf[:offset]), client)
}

func BuildRoomMessage(name, message string) []byte {
	buf := make([]byte, 256)
	offset := 0
	utils.WriteUint8(&buf, constant.ChatRoom, &offset)
	utils.WriteStringWithNull(&buf, []byte(name), &offset)
	utils.WriteStringWithNull(&buf, []byte(message), &offset)
	return buf[:offset]
}

func BuildChannelMessage(name, message string) []byte {
	buf := make([]byte, 256)
	offset := 0
	utils.WriteUint8(&buf, constant.ChatChannel, &offset)
	utils.WriteStringWithNull(&buf, []byte(name), &offset)
	utils.WriteStringWithNull(&buf, []byte(message), &offset)
	return buf[:offset]
}
