package room

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

type ChangeTeamService interface {
	Handle(ctx context.Context) error
}

type changeTeamServiceImpl struct {
	team   uint8
	client net.Conn
}

func NewChangeTeamService(p *packet.PacketData, client net.Conn) ChangeTeamService {
	return &changeTeamServiceImpl{
		team:   utils.ReadUint8(p.Data, &p.CurOffset),
		client: client,
	}
}

func (c *changeTeamServiceImpl) Handle(ctx context.Context) error {
	u := client.GetUserCacheClient().GetUserByConnection(ctx, c.client)
	if u == nil || u.CurrentRoomId == 0 {
		return errors.New("can't find user or user not in room")
	}

	room, err := client.GetRoomClient().GetRoomInfo(ctx, u.CurrentRoomId)
	if err != nil {
		return err
	}

	// 设置用户Team
	if err := client.GetUserCacheClient().SetUserRoom(ctx, u.UserID, room.RoomId, c.team); err != nil {
		return errors.New("set user room failed")
	}

	out_packet := out.BuildChangTeam(u.UserID, c.team)
	for i := range room.Users {
		tmp_u := client.GetUserCacheClient().GetUserByID(ctx, room.Users[i])
		if tmp_u == nil || tmp_u.CurrentRoomId == 0 {
			continue
		}

		rst := utils.BytesCombine(packet.BuildHeader(tmp_u.GetNextSeq(), constant.PacketTypeRoom), out_packet)
		packet.SendPacket(rst, tmp_u.CurrentConnection)
	}

	return nil
}
