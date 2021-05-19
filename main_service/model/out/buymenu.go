package out

import (
	"github.com/KouKouChan/YuriCore/main_service/model/user"
	"github.com/KouKouChan/YuriCore/utils"
)

func BuildUserBuymenu(buymenu *user.UserBuyMenu) []byte {
	buf := make([]byte, 512)
	offset := 0

	utils.WriteUint8(&buf, 0, &offset)

	utils.WriteUint8(&buf, 0, &offset)
	for i := range buymenu.PistolsTR {
		utils.WriteUint16(&buf, buymenu.PistolsTR[i], &offset)
	}

	utils.WriteUint8(&buf, 1, &offset)
	for i := range buymenu.PistolsTR {
		utils.WriteUint16(&buf, buymenu.ShotgunsTR[i], &offset)
	}

	utils.WriteUint8(&buf, 2, &offset)
	for i := range buymenu.PistolsTR {
		utils.WriteUint16(&buf, buymenu.SmgsTR[i], &offset)
	}

	utils.WriteUint8(&buf, 3, &offset)
	for i := range buymenu.PistolsTR {
		utils.WriteUint16(&buf, buymenu.RiflesTR[i], &offset)
	}

	utils.WriteUint8(&buf, 4, &offset)
	for i := range buymenu.PistolsTR {
		utils.WriteUint16(&buf, buymenu.MachinegunsTR[i], &offset)
	}

	utils.WriteUint8(&buf, 5, &offset)
	for i := range buymenu.PistolsTR {
		utils.WriteUint16(&buf, buymenu.EquipmentTR[i], &offset)
	}

	utils.WriteUint8(&buf, 6, &offset)
	for i := range buymenu.PistolsTR {
		utils.WriteUint16(&buf, buymenu.ClassesTR[i], &offset)
	}

	utils.WriteUint8(&buf, 7, &offset)
	for i := range buymenu.PistolsTR {
		utils.WriteUint16(&buf, buymenu.PistolsCT[i], &offset)
	}

	utils.WriteUint8(&buf, 8, &offset)
	for i := range buymenu.PistolsTR {
		utils.WriteUint16(&buf, buymenu.ShotgunsCT[i], &offset)
	}

	utils.WriteUint8(&buf, 9, &offset)
	for i := range buymenu.PistolsTR {
		utils.WriteUint16(&buf, buymenu.SmgsCT[i], &offset)
	}

	utils.WriteUint8(&buf, 10, &offset)
	for i := range buymenu.PistolsTR {
		utils.WriteUint16(&buf, buymenu.RiflesCT[i], &offset)
	}

	utils.WriteUint8(&buf, 11, &offset)
	for i := range buymenu.PistolsTR {
		utils.WriteUint16(&buf, buymenu.MachinegunsCT[i], &offset)
	}

	utils.WriteUint8(&buf, 12, &offset)
	for i := range buymenu.PistolsTR {
		utils.WriteUint16(&buf, buymenu.EquipmentCT[i], &offset)
	}

	utils.WriteUint8(&buf, 13, &offset)
	for i := range buymenu.PistolsTR {
		utils.WriteUint16(&buf, buymenu.ClassesCT[i], &offset)
	}

	utils.WriteUint8(&buf, 14, &offset)
	for i := range buymenu.PistolsTR {
		utils.WriteUint16(&buf, buymenu.MeleesTR[i], &offset)
	}

	utils.WriteUint8(&buf, 15, &offset)
	for i := range buymenu.PistolsTR {
		utils.WriteUint16(&buf, buymenu.MeleesCT[i], &offset)
	}
	return buf[:offset]
}
