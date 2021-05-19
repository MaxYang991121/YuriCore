package service

import (
	"context"

	"github.com/KouKouChan/YuriCore/main_service/model/server"
	"github.com/KouKouChan/YuriCore/room_service/client"
)

type RoomListService interface {
	Handle(ctx context.Context) ([]server.Room, error)
}

type roomListServiceImpl struct {
	serverID  uint8
	channelID uint8
}

func NewRoomListService(serverID, channelID uint8) RoomListService {
	return &roomListServiceImpl{
		serverID:  serverID,
		channelID: channelID,
	}
}

func (s *roomListServiceImpl) Handle(ctx context.Context) ([]server.Room, error) {
	return client.GetRoomTableClient().GetRoomList(ctx, s.serverID, s.channelID)
}
