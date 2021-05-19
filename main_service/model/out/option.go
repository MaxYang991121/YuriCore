package out

import (
	"github.com/KouKouChan/YuriCore/main_service/model/user"
	"github.com/KouKouChan/YuriCore/utils"
)

func BuildUserOptions(user *user.UserInfo) []byte {
	buf := make([]byte, 4096)
	offset := 0

	utils.WriteUint8(&buf, 0, &offset)
	utils.WriteLongString(&buf, user.Options, &offset)

	return buf[:offset]
}
