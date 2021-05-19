package client

import (
	"context"
	"sync"

	"github.com/KouKouChan/YuriCore/main_service/model/server"
)

type RoomTable interface {
	AddRoom(ctx context.Context, data *server.Room) (*server.Room, error)
	UpdateRoom(ctx context.Context, data *server.Room) error
	DeleteRoom(ctx context.Context, roomid uint16) error
	GetRoomList(ctx context.Context, serverID, ChannelID uint8) ([]server.Room, error)
	GetServerList(ctx context.Context) ([]server.Server, error)
	AddServer(ctx context.Context, srv server.Server)
	AddChannel(ctx context.Context, serverID uint8, chl server.Channel)
	GetRoomInfo(ctx context.Context, roomID uint16) (*server.Room, error)
	UpdateRoomSafe(ctx context.Context, data *server.Room) error
	AddUser(ctx context.Context, roomID uint16, userID uint32) (*server.Room, error)
	LeaveUser(ctx context.Context, roomID uint16, userID uint32) (*server.Room, error)
	SetUserHost(ctx context.Context, roomID uint16, userID uint32, name string) (*server.Room, error)
	StartGame(ctx context.Context, roomID uint16, userID uint32) (*server.Room, error)
	EndGame(ctx context.Context, roomID uint16, userID uint32) (*server.Room, error)
}

var (
	roomTableClient RoomTable
	roomTableOnce   sync.Once
)

func GetRoomTableClient() RoomTable {
	return roomTableClient
}

func InitRoomTableClient(client RoomTable) {
	roomTableOnce.Do(
		func() {
			roomTableClient = client
		},
	)
}
