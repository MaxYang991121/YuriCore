package out

import (
	"github.com/KouKouChan/YuriCore/main_service/model/user"
	"github.com/KouKouChan/YuriCore/utils"
)

func BuildUserCosmetics(cosmetics [5]user.UserCosmetics) []byte {
	buf := make([]byte, 512)
	offset := 0

	utils.WriteUint8(&buf, 1, &offset)

	for i := range cosmetics {
		utils.WriteUint8(&buf, uint8(i), &offset)
		utils.WriteStringWithNull(&buf, []byte(cosmetics[i].CosmeticsName), &offset)
		utils.WriteUint16(&buf, cosmetics[i].MainWeapon, &offset)
		utils.WriteUint16(&buf, cosmetics[i].MainBullet, &offset)
		utils.WriteUint16(&buf, cosmetics[i].SecondWeapon, &offset)
		utils.WriteUint16(&buf, cosmetics[i].SecondBullet, &offset)
		utils.WriteUint16(&buf, cosmetics[i].FlashbangNum, &offset)
		utils.WriteUint16(&buf, cosmetics[i].GrenadeID, &offset)
		utils.WriteUint16(&buf, cosmetics[i].SmokeNum, &offset)
		utils.WriteUint16(&buf, cosmetics[i].DefuserNum, &offset)
		utils.WriteUint16(&buf, cosmetics[i].TelescopeNum, &offset)
		utils.WriteUint16(&buf, cosmetics[i].BulletproofNum, &offset)
		utils.WriteUint16(&buf, cosmetics[i].KnifeID, &offset)
	}

	return buf[:offset]
}
