package controller

import (
	"context"
	"errors"

	"github.com/KouKouChan/YuriCore/main_service/model/server"
	"github.com/KouKouChan/YuriCore/room_service/service"
)

type UpdateRoomController interface {
	Handle(ctx context.Context) (*server.Room, error)
}

type updateRoomControllerImpl struct {
	room *server.Room
	safe bool
}

func NewUpdateRoomController(room *server.Room, safe bool) UpdateRoomController {
	return &updateRoomControllerImpl{
		room: room,
		safe: safe,
	}
}

func (u *updateRoomControllerImpl) Handle(ctx context.Context) (*server.Room, error) {
	if u.room.RoomId == 0 {
		return nil, errors.New("wrong room info")
	}

	if u.safe {
		return service.NewUpdateRoomService(u.room).UpdateRoomSafe(ctx)
	}

	return service.NewUpdateRoomService(u.room).UpdateRoom(ctx)
}
