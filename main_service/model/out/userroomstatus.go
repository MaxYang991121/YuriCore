package out

import "github.com/KouKouChan/YuriCore/utils"

func BuildUserReadyStatus(id uint32, status uint8) []byte {
	buf := make([]byte, 8)
	offset := 0
	utils.WriteUint8(&buf, 3, &offset)
	utils.WriteUint32(&buf, id, &offset)
	utils.WriteUint8(&buf, status, &offset)
	return buf[:offset]
}
