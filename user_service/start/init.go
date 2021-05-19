package start

import "github.com/KouKouChan/YuriCore/conf"

func Init(exePath string) {
	initConfig(exePath)
	if conf.Config.EnableDataBase == 1 {
		InitDBService()
	}
	initUserTable()
}
