package controller

import (
	"context"

	"github.com/KouKouChan/YuriCore/main_service/model/server"
	"github.com/KouKouChan/YuriCore/room_service/service"
)

type GetRoomInfoController interface {
	Handle(ctx context.Context) (*server.Room, error)
}

type getRoomInfoControllerImpl struct {
	roomID uint16
}

func NewGetRoomInfoController(roomID uint16) GetRoomInfoController {
	return &getRoomInfoControllerImpl{
		roomID: roomID,
	}
}

func (g *getRoomInfoControllerImpl) Handle(ctx context.Context) (*server.Room, error) {

	return service.NewGetRoomInfoService(g.roomID).Handle(ctx)
}
