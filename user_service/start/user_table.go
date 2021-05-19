package start

import (
	"github.com/KouKouChan/YuriCore/user_service/client"
	"github.com/KouKouChan/YuriCore/user_service/infrastructure/userTable"
)

func initUserTable() {
	client.InitUserTableClient(userTable.NewUserTableImpl())
}
