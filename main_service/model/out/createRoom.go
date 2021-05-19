package out

import (
	"github.com/KouKouChan/YuriCore/main_service/constant"
	"github.com/KouKouChan/YuriCore/main_service/model/server"
	"github.com/KouKouChan/YuriCore/main_service/model/user"
	"github.com/KouKouChan/YuriCore/utils"
)

func BuildCreateRoom(users []*user.UserInfo, caches []*user.UserCache, room *server.Room) []byte {
	if len(users) != len(caches) {
		return []byte{}
	}

	buf := make([]byte, 128)
	offset := 0

	utils.WriteUint8(&buf, constant.OUTCreateAndJoin, &offset)
	utils.WriteUint32(&buf, room.HostUserID, &offset)
	utils.WriteUint16(&buf, room.RoomId, &offset)
	utils.WriteUint8(&buf, 0x01, &offset)
	buf = utils.BytesCombine(buf[:offset], BuildRoomSetting(room, 0xFFFFFFFF7FFFFFFF))
	buf = append(buf, uint8(len(users)))
	for i := range users {
		if caches[i].UserID != users[i].UserID {
			continue
		}
		buf = utils.BytesCombine(buf, BuildUserNetInfo(caches[i]), BuildUserInfo(NewUserInfo(users[i]), false, 0xFFFFFFFF))
	}

	return buf
}
