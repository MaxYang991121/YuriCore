package out

import (
	"github.com/KouKouChan/YuriCore/main_service/constant"
	"github.com/KouKouChan/YuriCore/main_service/model/user"
	"github.com/KouKouChan/YuriCore/utils"
)

func BuildPlayerJoin(u *user.UserCache, info *user.UserInfo) []byte {
	buf := make([]byte, 1)
	offset := 0
	utils.WriteUint8(&buf, constant.OUTPlayerJoin, &offset)
	buf = utils.BytesCombine(buf, BuildUserNetInfo(u), BuildUserInfo(NewUserInfo(info), false, 0xFFFFFFFF))
	return buf
}
