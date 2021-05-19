package service

import (
	"context"

	"github.com/KouKouChan/YuriCore/main_service/model/server"
	"github.com/KouKouChan/YuriCore/room_service/client"
)

type StartService interface {
	Handle(ctx context.Context) (*server.Room, error)
}

type startServiceImpl struct {
	roomID uint16
	userID uint32
}

func NewStartService(roomID uint16, userID uint32) StartService {
	return &startServiceImpl{
		roomID: roomID,
		userID: userID,
	}
}

func (l *startServiceImpl) Handle(ctx context.Context) (*server.Room, error) {
	return client.GetRoomTableClient().StartGame(ctx, l.roomID, l.userID)
}
