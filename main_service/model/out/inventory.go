package out

import (
	"github.com/KouKouChan/YuriCore/main_service/model/user"
	"github.com/KouKouChan/YuriCore/utils"
)

func BuildUserInventory(inventory *user.Inventory) []byte {
	buf := make([]byte, 8192)
	offset := 0

	utils.WriteUint16(&buf, uint16(len(inventory.Items)), &offset)
	for i := range inventory.Items {
		utils.WriteUint16(&buf, uint16(i), &offset)
		utils.WriteUint8(&buf, inventory.Items[i].Existed, &offset) // 1 = existed
		if inventory.Items[i].Existed == 1 {
			utils.WriteUint16(&buf, inventory.Items[i].Id, &offset)
			utils.WriteUint16(&buf, inventory.Items[i].Count, &offset)
			utils.WriteUint8(&buf, 1, &offset)
			utils.WriteUint8(&buf, 0, &offset)
			utils.WriteUint32(&buf, 0, &offset)
			utils.WriteUint32(&buf, 0, &offset)
			utils.WriteUint32(&buf, 0, &offset)
		}
	}

	return buf[:offset]
}
