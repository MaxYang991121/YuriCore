package service

import (
	"context"

	"github.com/KouKouChan/YuriCore/main_service/model/server"
	"github.com/KouKouChan/YuriCore/room_service/client"
)

type LeaveService interface {
	Handle(ctx context.Context) (*server.Room, error)
}

type leaveServiceImpl struct {
	roomID uint16
	userID uint32
}

func NewLeaveService(roomID uint16, userID uint32) LeaveService {
	return &leaveServiceImpl{
		roomID: roomID,
		userID: userID,
	}
}

func (l *leaveServiceImpl) Handle(ctx context.Context) (*server.Room, error) {
	return client.GetRoomTableClient().LeaveUser(ctx, l.roomID, l.userID)
}
