package controller

import (
	"context"
	"errors"

	"github.com/KouKouChan/YuriCore/main_service/model/server"
	"github.com/KouKouChan/YuriCore/room_service/service"
)

type StartGameController interface {
	Handle(ctx context.Context) (*server.Room, error)
}

type startGameControllerImpl struct {
	roomID uint16
	userID uint32
}

func NewStartGameController(roomID uint16, userID uint32) StartGameController {
	return &startGameControllerImpl{
		roomID: roomID,
		userID: userID,
	}
}

func (l *startGameControllerImpl) Handle(ctx context.Context) (*server.Room, error) {
	if l.roomID == 0 || l.userID == 0 {
		return nil, errors.New("wrong userid or roomid")
	}

	return service.NewStartService(l.roomID, l.userID).Handle(ctx)
}
