package client

import (
	"context"
	"fmt"
	"sync"

	"github.com/KouKouChan/YuriCore/main_service/model/user"
)

type UserClient interface {
	Login(ctx context.Context, username, password string) (*user.UserInfo, int8)
	Register(ctx context.Context, username, password string) (bool, error)
	GetUserInfo(ctx context.Context, userID uint32) (*user.UserInfo, error)
	UserDown(ctx context.Context, userID uint32) (bool, error)
	GetUserFriends(ctx context.Context, userID uint32) ([]user.UserInfo, error)
	AddUserPoints(ctx context.Context, userID, add uint32) (uint32, error)
	AddUserCash(ctx context.Context, userID, add uint32) (uint32, error)
	UserPlayedGame(ctx context.Context, userID, IsWin, Kills, Deaths, HeadShots uint32) (*user.UserInfo, error)
	UserPayPoints(ctx context.Context, userID, used uint32) (uint32, error)
	UserPayCash(ctx context.Context, userID, used uint32) (uint32, error)
	UserAddItem(ctx context.Context, userID, item uint32) (*user.UserInfo, error)
	UserAddFriend(ctx context.Context, userID, friendID uint32) (*user.UserInfo, error)
	UpdateBag(ctx context.Context, UserID uint32, BagID uint16, Slot uint8, ItemID uint16) (*user.UserInfo, error)
	UpdateBuymenu(ctx context.Context, UserID uint32, BuymenuID uint16, Slot uint8, ItemID uint16) (*user.UserInfo, error)
	UpdateCosmetics(ctx context.Context, UserID uint32, CosmeticsID uint8, cosmetics *user.UserCosmetics) (*user.UserInfo, error)
	UpdateCampaign(ctx context.Context, UserID uint32, CampaignID uint8) (*user.UserInfo, error)
	UpdateOptions(ctx context.Context, UserID uint32, options []byte) (*user.UserInfo, error)
	UpdateNickname(ctx context.Context, UserID uint32, nickname string) (*user.UserInfo, error)
}

var (
	userClient UserClient
	userOnce   sync.Once
)

func GetUserClient() UserClient {
	return userClient
}

func InitUserClient(client UserClient) {
	userOnce.Do(
		func() {
			fmt.Println("User service connected")
			userClient = client
		},
	)
}
