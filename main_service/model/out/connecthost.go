package out

import (
	"github.com/KouKouChan/YuriCore/main_service/constant"
	"github.com/KouKouChan/YuriCore/utils"
)

func BuildConnectHost(ip uint32, port uint16) []byte {
	buf := make([]byte, 8)
	offset := 0
	utils.WriteUint8(&buf, constant.OUTConnectHost, &offset)
	utils.WriteUint32BE(&buf, ip, &offset)
	utils.WriteUint16(&buf, port, &offset)
	return buf[:offset]
}
