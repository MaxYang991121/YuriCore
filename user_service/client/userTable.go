package client

import (
	"context"
	"sync"

	"github.com/KouKouChan/YuriCore/main_service/model/user"
)

type UserTable interface {
	GetUserByID(ctx context.Context, id uint32) *user.UserInfo
	GetUserByUserName(ctx context.Context, username string) *user.UserInfo
	GetNewUserID(ctx context.Context) uint32
	UpdateUser(ctx context.Context, data *user.UserInfo) error
	DeleteUserByID(ctx context.Context, id uint32) error
	DeleteUserByName(ctx context.Context, username string) error
	AddUser(ctx context.Context, data *user.UserInfo) error
	UpdateBag(ctx context.Context, UserID uint32, BagID uint16, Slot uint8, ItemID uint16) (*user.UserInfo, error)
	UpdateBuymenu(ctx context.Context, UserID uint32, BuymenuID uint16, Slot uint8, ItemID uint16) (*user.UserInfo, error)
	UpdateCosmetics(ctx context.Context, UserID uint32, CosmeticsID uint8, cosmetics *user.UserCosmetics) (*user.UserInfo, error)
	UpdateCampaign(ctx context.Context, UserID uint32, CampaignID uint8) (*user.UserInfo, error)
	GetUserFriends(ctx context.Context, UserID uint32) ([]user.UserInfo, error)
	AddUserPoints(ctx context.Context, UserID uint32, num uint64) (uint64, error)
	AddUserCash(ctx context.Context, UserID uint32, num uint64) (uint64, error)
	PayPoints(ctx context.Context, UserID uint32, num uint64) (uint64, error)
	PayCash(ctx context.Context, UserID uint32, num uint64) (uint64, error)
	AddFriend(ctx context.Context, UserID uint32, friend string) (*user.UserInfo, error)
	UpdateOption(ctx context.Context, UserID uint32, Options []byte) (*user.UserInfo, error)
	UpdateNickname(ctx context.Context, UserID uint32, nickname string) (*user.UserInfo, error)
}

var (
	userTableClient UserTable
	userTableOnce   sync.Once
)

func GetUserTableClient() UserTable {
	return userTableClient
}

func InitUserTableClient(client UserTable) {
	userTableOnce.Do(
		func() {
			userTableClient = client
		},
	)
}
