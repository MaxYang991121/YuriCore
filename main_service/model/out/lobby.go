package out

import (
	"github.com/KouKouChan/YuriCore/main_service/model/user"
	"github.com/KouKouChan/YuriCore/utils"
)

func BuildLobbyReply(users []user.UserInfo) []byte {
	buf := make([]byte, 3)
	offset := 0

	utils.WriteUint8(&buf, 0, &offset)
	utils.WriteUint16(&buf, uint16(len(users)), &offset)
	for i := range users {
		head := make([]byte, 64)
		head_offset := 0
		utils.WriteUint32(&head, users[i].UserID, &head_offset)
		utils.WriteStringWithNull(&head, []byte(users[i].NickName), &head_offset)
		buf = utils.BytesCombine(buf, head[:head_offset])
		buf = utils.BytesCombine(buf, BuildUserInfo(NewUserInfo(&users[i]), false, 0xFFFFFFFF))
	}
	return buf
}

func BuildLobbyLeave(userID uint32) []byte {
	buf := make([]byte, 8)
	offset := 0

	utils.WriteUint8(&buf, 2, &offset)
	utils.WriteUint32(&buf, userID, &offset)
	return buf[:offset]
}

func BuildLobbyJoin(user *user.UserInfo) []byte {
	buf := make([]byte, 128)
	offset := 0

	utils.WriteUint8(&buf, 1, &offset)
	utils.WriteUint32(&buf, user.UserID, &offset)
	utils.WriteStringWithNull(&buf, []byte(user.NickName), &offset)
	buf = utils.BytesCombine(buf[:offset], BuildUserInfo(NewUserInfo(user), false, 0xFFFFFFFF))
	return buf
}
