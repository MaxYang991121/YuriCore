package out

import (
	"github.com/KouKouChan/YuriCore/main_service/constant"
	"github.com/KouKouChan/YuriCore/utils"
)

func BuildChangTeam(id uint32, team uint8) []byte {
	buf := make([]byte, 7)
	offset := 0
	utils.WriteUint8(&buf, constant.OUTsetUserTeam, &offset)
	utils.WriteUint32(&buf, id, &offset)
	utils.WriteUint8(&buf, team, &offset)
	return buf[:offset]
}
