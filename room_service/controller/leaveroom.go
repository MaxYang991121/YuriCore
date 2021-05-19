package controller

import (
	"context"
	"errors"

	"github.com/KouKouChan/YuriCore/main_service/model/server"
	"github.com/KouKouChan/YuriCore/room_service/service"
)

type LeaveRoomInfoController interface {
	Handle(ctx context.Context) (*server.Room, error)
}

type leaveRoomInfoControllerImpl struct {
	roomID uint16
	userID uint32
}

func NewLeaveRoomInfoController(roomID uint16, userID uint32) LeaveRoomInfoController {
	return &leaveRoomInfoControllerImpl{
		roomID: roomID,
		userID: userID,
	}
}

func (l *leaveRoomInfoControllerImpl) Handle(ctx context.Context) (*server.Room, error) {
	if l.roomID == 0 || l.userID == 0 {
		return nil, errors.New("wrong userid or roomid")
	}

	return service.NewLeaveService(l.roomID, l.userID).Handle(ctx)
}
