package controller

import (
	"context"
	"errors"

	"github.com/KouKouChan/YuriCore/main_service/model/server"
	"github.com/KouKouChan/YuriCore/room_service/service"
)

type JoinRoomInfoController interface {
	Handle(ctx context.Context) (*server.Room, error)
}

type joinRoomInfoControllerImpl struct {
	roomID uint16
	userID uint32
}

func NewJoinRoomInfoController(roomID uint16, userID uint32) JoinRoomInfoController {
	return &joinRoomInfoControllerImpl{
		roomID: roomID,
		userID: userID,
	}
}

func (l *joinRoomInfoControllerImpl) Handle(ctx context.Context) (*server.Room, error) {
	if l.roomID == 0 || l.userID == 0 {
		return nil, errors.New("wrong userid or roomid")
	}

	return service.NewJoinService(l.roomID, l.userID).Handle(ctx)
}
