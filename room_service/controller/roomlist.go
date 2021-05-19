package controller

import (
	"context"
	"errors"

	"github.com/KouKouChan/YuriCore/main_service/model/server"
	"github.com/KouKouChan/YuriCore/room_service/service"
)

type RoomListController interface {
	Handle(ctx context.Context) ([]server.Room, error)
}

type roomListControllerImpl struct {
	serverID  uint8
	channelID uint8
}

func NewRoomListController(serverID, channelID uint8) RoomListController {
	return &roomListControllerImpl{
		serverID:  serverID,
		channelID: channelID,
	}
}

func (r *roomListControllerImpl) Handle(ctx context.Context) ([]server.Room, error) {
	if r.channelID == 0 || r.serverID == 0 {
		return nil, errors.New("wrong channelID or serverID")
	}

	return service.NewRoomListService(r.serverID, r.channelID).Handle(ctx)
}
