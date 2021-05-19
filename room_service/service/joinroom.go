package service

import (
	"context"

	"github.com/KouKouChan/YuriCore/main_service/model/server"
	"github.com/KouKouChan/YuriCore/room_service/client"
)

type JoinService interface {
	Handle(ctx context.Context) (*server.Room, error)
}

type joinServiceImpl struct {
	roomID uint16
	userID uint32
}

func NewJoinService(roomID uint16, userID uint32) JoinService {
	return &joinServiceImpl{
		roomID: roomID,
		userID: userID,
	}
}

func (l *joinServiceImpl) Handle(ctx context.Context) (*server.Room, error) {
	return client.GetRoomTableClient().AddUser(ctx, l.roomID, l.userID)
}
