package service

import (
	"context"

	"github.com/KouKouChan/YuriCore/main_service/model/server"
	"github.com/KouKouChan/YuriCore/room_service/client"
)

type GetRoomInfoService interface {
	Handle(ctx context.Context) (*server.Room, error)
}

type getRoomInfoServiceImpl struct {
	roomID uint16
}

func NewGetRoomInfoService(roomID uint16) GetRoomInfoService {
	return &getRoomInfoServiceImpl{
		roomID: roomID,
	}
}

func (g *getRoomInfoServiceImpl) Handle(ctx context.Context) (*server.Room, error) {

	return client.GetRoomTableClient().GetRoomInfo(ctx, g.roomID)
}
