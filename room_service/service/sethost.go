package service

import (
	"context"

	"github.com/KouKouChan/YuriCore/main_service/model/server"
	"github.com/KouKouChan/YuriCore/room_service/client"
)

type SetHostService interface {
	Handle(ctx context.Context) (*server.Room, error)
}

type setHostServiceImpl struct {
	roomID uint16
	userID uint32
	name   string
}

func NewSetHostService(roomID uint16, userID uint32, name string) SetHostService {
	return &setHostServiceImpl{
		roomID: roomID,
		userID: userID,
		name:   name,
	}
}

func (s *setHostServiceImpl) Handle(ctx context.Context) (*server.Room, error) {

	return client.GetRoomTableClient().SetUserHost(ctx, s.roomID, s.userID, s.name)
}
