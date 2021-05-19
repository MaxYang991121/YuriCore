package start

import (
	"fmt"

	"github.com/KouKouChan/YuriCore/conf"
	"github.com/KouKouChan/YuriCore/main_service/client"
	"github.com/KouKouChan/YuriCore/main_service/infrastructure/user_service"
)

func initUserService() {
	client.InitUserClient(user_service.NewUserService(fmt.Sprintf("%s:%d", conf.Config.UserAdress, conf.Config.UserPort)))
}
