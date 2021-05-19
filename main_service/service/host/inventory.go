package host

import (
	"context"
	"errors"
	"net"

	"github.com/KouKouChan/YuriCore/main_service/client"
	"github.com/KouKouChan/YuriCore/main_service/constant"
	"github.com/KouKouChan/YuriCore/main_service/model/out"
	"github.com/KouKouChan/YuriCore/main_service/model/packet"
	"github.com/KouKouChan/YuriCore/utils"
)

type GetInventoryService interface {
	Handle(ctx context.Context) error
}

type getInventoryServiceImpl struct {
	dest   uint32
	client net.Conn
}

func NewGetInventoryService(p *packet.PacketData, client net.Conn) GetInventoryService {
	return &getInventoryServiceImpl{
		dest:   utils.ReadUint32(p.Data, &p.CurOffset),
		client: client,
	}
}

func (g *getInventoryServiceImpl) Handle(ctx context.Context) error {
	u := client.GetUserCacheClient().GetUserByConnection(ctx, g.client)
	if u == nil || u.CurrentRoomId == 0 {
		return errors.New("can't find user or user not in room")
	}

	room, err := client.GetRoomClient().GetRoomInfo(ctx, u.CurrentRoomId)
	if err != nil {
		return err
	}
	if room.HostUserID != u.UserID {
		return errors.New("user try to get other player's inventory but is not host")
	}

	dest_player := client.GetUserCacheClient().GetUserByID(ctx, g.dest)
	if dest_player == nil || dest_player.CurrentRoomId != u.CurrentRoomId {
		return errors.New("can't find dest player or player not in room")
	}

	rst := utils.BytesCombine(packet.BuildHeader(u.GetNextSeq(), constant.PacketTypeHost), out.BuildSetUserInventory(dest_player))
	packet.SendPacket(rst, u.CurrentConnection)

	return nil
}
