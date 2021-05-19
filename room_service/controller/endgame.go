package controller

import (
	"context"
	"errors"

	"github.com/KouKouChan/YuriCore/main_service/model/server"
	"github.com/KouKouChan/YuriCore/room_service/service"
)

type EndGameController interface {
	Handle(ctx context.Context) (*server.Room, error)
}

type endGameControllerImpl struct {
	roomID uint16
	userID uint32
}

func NewEndGameController(roomID uint16, userID uint32) EndGameController {
	return &endGameControllerImpl{
		roomID: roomID,
		userID: userID,
	}
}

func (l *endGameControllerImpl) Handle(ctx context.Context) (*server.Room, error) {
	if l.roomID == 0 || l.userID == 0 {
		return nil, errors.New("wrong userid or roomid")
	}

	return service.NewEndService(l.roomID, l.userID).Handle(ctx)
}
