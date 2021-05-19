package out

import "github.com/KouKouChan/YuriCore/utils"

func BuildUserStart(id uint32, username, nickname []byte) []byte {
	userbuf := make([]byte, 128)
	offset := 0
	utils.WriteUint32(&userbuf, id, &offset)
	utils.WriteStringWithNull(&userbuf, username, &offset)
	utils.WriteStringWithNull(&userbuf, nickname, &offset)
	utils.WriteUint8(&userbuf, 0, &offset)
	utils.WriteUint8(&userbuf, 0, &offset)
	utils.WriteUint8(&userbuf, 0, &offset)

	return userbuf[:offset]
}
