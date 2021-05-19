package service

import (
	"context"

	"github.com/KouKouChan/YuriCore/main_service/model/server"
	"github.com/KouKouChan/YuriCore/room_service/client"
)

type NewRoomService interface {
	Handle(ctx context.Context) (*server.Room, error)
}

type newRoomServiceImpl struct {
	room server.Room
}

func NewNewRoomService(room server.Room) NewRoomService {
	return &newRoomServiceImpl{
		room: room,
	}
}

func (n *newRoomServiceImpl) Handle(ctx context.Context) (*server.Room, error) {

	return client.GetRoomTableClient().AddRoom(ctx, &n.room)
}
