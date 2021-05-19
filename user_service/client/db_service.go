package client

import (
	"context"
	"fmt"
	"sync"

	"github.com/KouKouChan/YuriCore/main_service/model/user"
)

type DB interface {
	GetUser(ctx context.Context, username string) (*user.UserInfo, error)
	UpdateUser(ctx context.Context, user *user.UserInfo) error
	GetUserByNickName(ctx context.Context, nickname string) (*user.UserInfo, error)
}

var (
	dbClient DB
	dbOnce   sync.Once
)

func GetDBClient() DB {
	return dbClient
}

func InitDBClient(client DB) {
	dbOnce.Do(
		func() {
			fmt.Println("DB service connected")
			dbClient = client
		},
	)
}
