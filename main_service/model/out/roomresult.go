package out

import (
	"github.com/KouKouChan/YuriCore/main_service/constant"
	"github.com/KouKouChan/YuriCore/main_service/model/server"
	"github.com/KouKouChan/YuriCore/utils"
)

func BuildRoomResult(room *server.Room) []byte {
	buf := make([]byte, 512)
	offset := 0
	utils.WriteUint8(&buf, constant.OUTSetGameResult, &offset)
	utils.WriteUint8(&buf, 0, &offset)
	utils.WriteUint8(&buf, 0, &offset)
	utils.WriteUint8(&buf, 0, &offset)
	utils.WriteUint8(&buf, 0, &offset)
	utils.WriteUint8(&buf, 0, &offset)
	utils.WriteUint8(&buf, 0, &offset)

	return buf[:offset]
}
