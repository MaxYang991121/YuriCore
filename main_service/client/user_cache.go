package client

import (
	"context"
	"net"
	"sync"

	"github.com/KouKouChan/YuriCore/main_service/model/user"
)

type UserCache interface {
	GetUserByID(ctx context.Context, id uint32) *user.UserCache
	GetUserByUserName(ctx context.Context, username string) *user.UserCache
	GetUserByConnection(ctx context.Context, client net.Conn) *user.UserCache
	GetChannelUsers(ctx context.Context, serverID, channelID uint8) []uint32
	DeleteUserByID(ctx context.Context, id uint32)
	DeleteUserByName(ctx context.Context, username string)
	DeleteUserByConnection(ctx context.Context, client net.Conn)
	SetUser(ctx context.Context, data *user.UserCache) error
	SetUserChannel(ctx context.Context, userID uint32, serverID, channelID uint8) error
	SetUserQuitChannel(ctx context.Context, userID uint32) error
	SetUserRoom(ctx context.Context, userID uint32, roomID uint16, team uint8) error
	SetUserStatus(ctx context.Context, userID uint32, status uint8) error
	QuitUserRoom(ctx context.Context, userID uint32) error
	FlushUserInventory(ctx context.Context, userID uint32, inventory *user.Inventory) error
	FlushUserUDP(ctx context.Context, userID uint32, portId uint16, localPort uint16, externalPort uint16, externalIPAddress, localIpAddress uint32) (uint16, error)
	SetUserIngame(ctx context.Context, userID uint32, ingame bool) error
	ResetKillNum(ctx context.Context, userID uint32) error
	ResetDeadNum(ctx context.Context, userID uint32) error
	ResetAssistNum(ctx context.Context, userID uint32) error
	GetChannelNoRoomUsers(ctx context.Context, serverID, channelID uint8) []uint32
	FlushUserRoomData(ctx context.Context, userID uint32, data []byte) error
	SetNickname(ctx context.Context, userID uint32, nickname string) error
}

var (
	userCacheClient UserCache
	userCacheOnce   sync.Once
)

func GetUserCacheClient() UserCache {
	return userCacheClient
}

func InitUserCacheClient(client UserCache) {
	userCacheOnce.Do(
		func() {
			userCacheClient = client
		},
	)
}
