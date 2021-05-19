package out

import (
	"github.com/KouKouChan/YuriCore/main_service/constant"
	"github.com/KouKouChan/YuriCore/utils"
)

func BuildStartRoom(id uint32) []byte {
	buf := make([]byte, 5)
	offset := 0
	utils.WriteUint8(&buf, constant.OUTStartRoom, &offset)
	utils.WriteUint32(&buf, id, &offset)
	return buf[:offset]
}
