package controller

import (
	"context"
	"errors"

	"github.com/KouKouChan/YuriCore/main_service/model/server"
	"github.com/KouKouChan/YuriCore/room_service/service"
)

type SetHostController interface {
	Handle(ctx context.Context) (*server.Room, error)
}

type setHostControllerImpl struct {
	roomID uint16
	userID uint32
	name   string
}

func NewSetHostController(roomID uint16, userID uint32, name string) SetHostController {
	return &setHostControllerImpl{
		roomID: roomID,
		userID: userID,
		name:   name,
	}
}

func (s *setHostControllerImpl) Handle(ctx context.Context) (*server.Room, error) {
	if s.roomID == 0 || s.userID == 0 {
		return nil, errors.New("wrong roomid or userid")
	}

	return service.NewSetHostService(s.roomID, s.userID, s.name).Handle(ctx)
}
