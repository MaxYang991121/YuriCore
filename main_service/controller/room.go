package controller

import (
	"context"
	"fmt"
	"net"

	"github.com/KouKouChan/YuriCore/main_service/model/packet"
	"github.com/KouKouChan/YuriCore/main_service/service/room"
	"github.com/KouKouChan/YuriCore/utils"
)

type RoomController interface {
	Handle(ctx context.Context) error
}

type roomControllerImpl struct {
	roomType uint8
	client   net.Conn
	packet   *packet.PacketData
}

func GetRoomController(p *packet.PacketData, client net.Conn) RoomController {
	roomlist := roomControllerImpl{}

	roomlist.roomType = utils.ReadUint8(p.Data, &p.CurOffset)
	roomlist.client = client
	roomlist.packet = p

	return &roomlist
}

func (r *roomControllerImpl) Handle(ctx context.Context) error {

	switch r.roomType {
	case 0:
		return room.NewCreateRoomService(r.packet, r.client).Handle(ctx)
	case 1:
		return room.NewJoinRoomService(r.packet, r.client).Handle(ctx)
	case 2:
		return room.NewLeaveRoomService(r.client).Handle(ctx)
	case 3:
		return room.NewReadyService(r.packet, r.client).Handle(ctx)
	case 4:
		return room.NewStartService(r.client).Handle(ctx)
	case 5:
		return room.NewUpdateRoomService(r.packet, r.client).Handle(ctx)
	case 6:
		return room.NewCloseResultService(r.client).Handle(ctx)
	case 17:
		return room.NewChangeTeamService(r.packet, r.client).Handle(ctx)
	default:
		return fmt.Errorf("Unknown room packet %+v", r.roomType)
	}
}
