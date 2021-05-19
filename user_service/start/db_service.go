package start

import (
	"github.com/KouKouChan/YuriCore/conf"
	"github.com/KouKouChan/YuriCore/user_service/client"
	"github.com/KouKouChan/YuriCore/user_service/infrastructure/db_service"
)

func InitDBService() {
	client.InitDBClient(db_service.NewDBImpl(
		conf.Config.DBUserName,
		conf.Config.DBpassword,
		conf.Config.DBaddress,
		conf.Config.DBport,
	))
}
