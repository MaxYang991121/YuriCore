package start

import (
	"math"

	"github.com/KouKouChan/YuriCore/conf"
	"github.com/KouKouChan/YuriCore/verbose"
)

func initConfig(exePath string) {
	conf.Config.InitConf(exePath + "/server.conf")
	if conf.Config.MaxUsers <= 0 {
		conf.Config.MaxUsers = math.MaxUint32
	}
	verbose.Level = conf.Config.DebugLevel
	verbose.LogFile = conf.Config.LogFile
	if conf.Config.LogFile != 0 {
		verbose.InitLoger(exePath, "roomservice.log")
	}

}
