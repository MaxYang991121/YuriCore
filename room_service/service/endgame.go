package service

import (
	"context"

	"github.com/KouKouChan/YuriCore/main_service/model/server"
	"github.com/KouKouChan/YuriCore/room_service/client"
)

type EndService interface {
	Handle(ctx context.Context) (*server.Room, error)
}

type endServiceImpl struct {
	roomID uint16
	userID uint32
}

func NewEndService(roomID uint16, userID uint32) EndService {
	return &endServiceImpl{
		roomID: roomID,
		userID: userID,
	}
}

func (l *endServiceImpl) Handle(ctx context.Context) (*server.Room, error) {
	return client.GetRoomTableClient().EndGame(ctx, l.roomID, l.userID)
}
