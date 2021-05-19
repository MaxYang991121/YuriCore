package out

import (
	"github.com/KouKouChan/YuriCore/main_service/model/user"
	"github.com/KouKouChan/YuriCore/utils"
)

func BuildUserBag(bags [3]user.UserLoadout) []byte {
	buf := make([]byte, 256)
	offset := 0

	utils.WriteUint8(&buf, 2, &offset)
	utils.WriteUint16(&buf, 0, &offset)
	utils.WriteUint8(&buf, 0, &offset)
	utils.WriteUint8(&buf, uint8(len(bags)), &offset)
	utils.WriteUint8(&buf, 4, &offset)

	utils.WriteUint16(&buf, bags[0].MainWeapon, &offset)
	utils.WriteUint16(&buf, bags[0].SecondWeapon, &offset)
	utils.WriteUint16(&buf, bags[0].Knife, &offset)
	utils.WriteUint16(&buf, bags[0].Grenade, &offset)

	utils.WriteUint16(&buf, bags[1].MainWeapon, &offset)
	utils.WriteUint16(&buf, bags[1].SecondWeapon, &offset)
	utils.WriteUint16(&buf, bags[1].Knife, &offset)
	utils.WriteUint16(&buf, bags[1].Grenade, &offset)

	utils.WriteUint16(&buf, bags[2].MainWeapon, &offset)
	utils.WriteUint16(&buf, bags[2].SecondWeapon, &offset)
	utils.WriteUint16(&buf, bags[2].Knife, &offset)
	utils.WriteUint16(&buf, bags[2].Grenade, &offset)
	return buf[:offset]
}
