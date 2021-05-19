package start

import (
	"github.com/KouKouChan/YuriCore/conf"
	"github.com/KouKouChan/YuriCore/utils"
)

func Init(exePath string) {
	initConfig(exePath)
	utils.InitConverter(conf.Config.CodePage)
	initUserService()
	initUserCache()
	initRoomService()
}
