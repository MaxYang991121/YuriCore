package out

import (
	"github.com/KouKouChan/YuriCore/main_service/constant"
	"github.com/KouKouChan/YuriCore/main_service/model/user"
	"github.com/KouKouChan/YuriCore/utils"
)

func BuildSetUserInventory(u *user.UserCache) []byte {
	buf := make([]byte, 8+8*len(u.UserInventory.Items))
	offset := 0
	utils.WriteUint8(&buf, constant.OUTSetInventory, &offset)
	utils.WriteUint32(&buf, u.UserID, &offset)
	utils.WriteUint16(&buf, uint16(len(u.UserInventory.Items)), &offset)
	for _, v := range u.UserInventory.Items {
		utils.WriteUint16(&buf, v.Id, &offset)
		utils.WriteUint16(&buf, v.Count, &offset)
		utils.WriteUint32(&buf, v.Time, &offset)
	}
	return buf[:offset]
}
