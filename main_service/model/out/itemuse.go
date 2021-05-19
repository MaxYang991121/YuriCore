package out

import (
	"github.com/KouKouChan/YuriCore/main_service/constant"
	"github.com/KouKouChan/YuriCore/utils"
)

func BuildHostItemUsing(userID uint32, itemID uint16, num uint8) []byte {
	buf := make([]byte, 32)
	offset := 0

	utils.WriteUint8(&buf, constant.OUTItemUsing, &offset)
	utils.WriteUint32(&buf, userID, &offset)
	utils.WriteUint16(&buf, itemID, &offset)
	utils.WriteUint8(&buf, num, &offset)

	return buf[:offset]
}
