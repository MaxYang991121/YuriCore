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
	. "github.com/KouKouChan/YuriCore/verbose"
)

type ReadyService interface {
	Handle(ctx context.Context) error
}

type readyServiceImpl struct {
	Unk00  uint8
	client net.Conn
}

func NewReadyService(p *packet.PacketData, client net.Conn) ReadyService {
	return &readyServiceImpl{
		Unk00:  utils.ReadUint8(p.Data, &p.CurOffset),
		client: client,
	}
}

func (c *readyServiceImpl) Handle(ctx context.Context) error {
	u := client.GetUserCacheClient().GetUserByConnection(ctx, c.client)
	if u == nil || u.CurrentRoomId == 0 {
		return errors.New("can't find user or user not in room")
	}

	room, err := client.GetRoomClient().GetRoomInfo(ctx, u.CurrentRoomId)
	if err != nil {
		return err
	}

	//设置新的状态
	if u.Currentstatus == constant.UserReady {
		if err = client.GetUserCacheClient().SetUserStatus(ctx, u.UserID, constant.UserNotReady); err != nil {
			return nil
		}
		DebugInfo(1, "User", u.UserName, "unreadied in room", room.RoomName, "id", room.RoomId)
	} else {
		if err = client.GetUserCacheClient().SetUserStatus(ctx, u.UserID, constant.UserReady); err != nil {
			return nil
		}
		DebugInfo(1, "User", u.UserName, "readied in room", room.RoomName, "id", room.RoomId)
	}

	// 发送数据包
	ustatus := out.BuildUserReadyStatus(u.UserID, u.Currentstatus)
	for i := range room.Users {
		dest_player := client.GetUserCacheClient().GetUserByID(ctx, room.Users[i])
		if dest_player == nil {
			continue
		}
		rst := utils.BytesCombine(packet.BuildHeader(dest_player.GetNextSeq(), constant.PacketTypeRoom), ustatus)
		packet.SendPacket(rst, dest_player.CurrentConnection)
	}
	return nil
}
