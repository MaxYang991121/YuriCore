package controller

import (
	"context"
	"errors"
	"net"

	"github.com/KouKouChan/YuriCore/main_service/model/packet"
	"github.com/KouKouChan/YuriCore/main_service/service/roomlist"
	"github.com/KouKouChan/YuriCore/utils"
)

type RoomListController interface {
	Handle(ctx context.Context) error
}

type roomListControllerImpl struct {
	serverIndex  uint8
	channelIndex uint8
	client       net.Conn
}

func GetRoomListController(p *packet.PacketData, client net.Conn) RoomListController {
	roomlist := roomListControllerImpl{}

	roomlist.serverIndex = utils.ReadUint8(p.Data, &p.CurOffset)
	roomlist.channelIndex = utils.ReadUint8(p.Data, &p.CurOffset)
	roomlist.client = client

	return &roomlist
}

func (r *roomListControllerImpl) Handle(ctx context.Context) error {
	if r.channelIndex == 0 || r.serverIndex == 0 {
		return errors.New("wrong channelID or serverID")
	}

	return roomlist.NewRoomListService(r.serverIndex, r.channelIndex, r.client).Handler(ctx)
}
