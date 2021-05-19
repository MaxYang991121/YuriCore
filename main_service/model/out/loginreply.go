package out

import (
	"github.com/KouKouChan/YuriCore/main_service/constant"
	"github.com/KouKouChan/YuriCore/utils"
)

func BuildLoginReply() []byte {
	buf := make([]byte, 32)
	offset := 0
	utils.WriteUint8(&buf, 0, &offset)
	utils.WriteStringWithNull(&buf, []byte(constant.ReplyYes), &offset)
	utils.WriteUint8(&buf, 0, &offset)
	return buf[:offset]
}

func BuildNicknameReply() []byte {
	buf := make([]byte, 32)
	offset := 0
	utils.WriteUint8(&buf, 1, &offset)
	utils.WriteStringWithNull(&buf, []byte(constant.ReplyYes), &offset)
	utils.WriteUint8(&buf, 0, &offset)
	return buf[:offset]
}
