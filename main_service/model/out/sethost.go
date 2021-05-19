package out

import (
	"github.com/KouKouChan/YuriCore/main_service/constant"
	"github.com/KouKouChan/YuriCore/utils"
)

func BuildSetHost(id uint32, isHost uint8) []byte {
	buf := make([]byte, 6)
	offset := 0
	utils.WriteUint8(&buf, constant.OUTSetHost, &offset)
	utils.WriteUint32(&buf, id, &offset)
	utils.WriteUint8(&buf, isHost, &offset)
	return buf[:offset]
}
