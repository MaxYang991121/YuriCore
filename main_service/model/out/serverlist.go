package out

import (
	"github.com/KouKouChan/YuriCore/main_service/model/server"
	"github.com/KouKouChan/YuriCore/utils"
)

func BuildServerList(servers []server.Server) []byte {
	buf := make([]byte, 4096)
	offset := 0
	utils.WriteUint8(&buf, uint8(len(servers)), &offset)
	for i := range servers {
		utils.WriteUint8(&buf, servers[i].ServerIndex, &offset)
		utils.WriteUint8(&buf, servers[i].ServerStatus, &offset)
		utils.WriteUint8(&buf, servers[i].ServerType, &offset)
		ansiName, _ := utils.Utf8ToLocal(servers[i].ServerName)
		utils.WriteStringWithNull(&buf, []byte(ansiName), &offset)
		utils.WriteUint8(&buf, uint8(len(servers[i].Channels)), &offset)
		for j := range servers[i].Channels {
			utils.WriteUint8(&buf, servers[i].Channels[j].ChannelIndex, &offset)
			ansiName, _ := utils.Utf8ToLocal(servers[i].Channels[j].ChannelName)
			utils.WriteStringWithNull(&buf, []byte(ansiName), &offset)
			utils.WriteUint16(&buf, 0x00, &offset)
		}
	}
	return buf[:offset]
}
