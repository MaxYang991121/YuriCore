package client

import (
	"context"
	"fmt"
	"sync"

	"github.com/KouKouChan/YuriCore/main_service/model/server"
)

type RoomClient interface {
	GetServiceList(ctx context.Context) ([]server.Server, error)
	GetRoomList(ctx context.Context, ServerIndex, ChannelIndex uint8) ([]server.Room, error)
	NewRoom(ctx context.Context, room *server.Room) (*server.Room, error)
	UpdateRoom(ctx context.Context, room *server.Room) (*server.Room, error)
	JoinRoom(ctx context.Context, userID uint32, roomID uint16) (*server.Room, error)
	LeaveRoom(ctx context.Context, userID uint32, roomID uint16) (*server.Room, error)
	StartGame(ctx context.Context, userID uint32, roomID uint16) (*server.Room, error)
	GetRoomInfo(ctx context.Context, roomID uint16) (*server.Room, error)
	UpdateRoomSafe(ctx context.Context, room *server.Room) (*server.Room, error)
	SetRoomHost(ctx context.Context, userID uint32, name string, roomID uint16) (*server.Room, error)
	DelRoom(ctx context.Context, roomID uint16) error
	EndGame(ctx context.Context, userID uint32, roomID uint16) (*server.Room, error)
}

var (
	roomClient RoomClient
	roomOnce   sync.Once
)

func GetRoomClient() RoomClient {
	return roomClient
}

func InitRoomClient(client RoomClient) {
	roomOnce.Do(
		func() {
			fmt.Println("Room service connected")
			roomClient = client
		},
	)
}
