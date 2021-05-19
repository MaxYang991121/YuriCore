package out

import (
	"github.com/KouKouChan/YuriCore/main_service/model/user"
	"github.com/KouKouChan/YuriCore/utils"
)

func BuildUserNetInfo(u *user.UserCache) []byte {
	buf := make([]byte, 29)
	offset := 0
	utils.WriteUint32(&buf, u.UserID, &offset)
	utils.WriteUint8(&buf, u.CurrentTeam, &offset)
	utils.WriteUint8(&buf, u.Currentstatus, &offset)
	utils.WriteUint8(&buf, 0, &offset)                              // status
	utils.WriteUint32BE(&buf, u.NetInfo.ExternalIpAddress, &offset) //externalIpAddress
	utils.WriteUint16(&buf, u.NetInfo.ExternalServerPort, &offset)  //externalServerPort
	utils.WriteUint16(&buf, u.NetInfo.ExternalClientPort, &offset)  //externalClientPort
	utils.WriteUint32BE(&buf, u.NetInfo.LocalIpAddress, &offset)    //localIpAddress
	utils.WriteUint16(&buf, u.NetInfo.LocalServerPort, &offset)     //localServerPort
	utils.WriteUint16(&buf, u.NetInfo.LocalClientPort, &offset)     //localClientPort
	return buf[:offset]
}
