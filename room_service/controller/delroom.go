package controller

import (
	"context"
	"errors"

	"github.com/KouKouChan/YuriCore/room_service/service"
)

type DelRoomController interface {
	Handle(ctx context.Context) error
}

type delRoomControllerImpl struct {
	roomID uint16
}

func NewDelRoomController(roomID uint16) DelRoomController {
	return &delRoomControllerImpl{
		roomID: roomID,
	}
}

func (d *delRoomControllerImpl) Handle(ctx context.Context) error {
	if d.roomID == 0 {
		return errors.New("wrong roomid")
	}

	return service.NewDelRoomService(d.roomID).Handle(ctx)
}
