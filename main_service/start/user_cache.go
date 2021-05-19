package start

import (
	"github.com/KouKouChan/YuriCore/main_service/client"
	"github.com/KouKouChan/YuriCore/main_service/infrastructure/user_cache"
)

func initUserCache() {
	client.InitUserCacheClient(user_cache.NewUserCacheImpl())
}
