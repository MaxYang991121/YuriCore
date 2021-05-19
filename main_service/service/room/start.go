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

type StartService interface {
	Handle(ctx context.Context) error
}

type startServiceImpl struct {
	client net.Conn
}

func NewStartService(client net.Conn) StartService {
	return &startServiceImpl{
		client: client,
	}
}

func (c *startServiceImpl) Handle(ctx context.Context) error {
	u := client.GetUserCacheClient().GetUserByConnection(ctx, c.client)
	if u == nil {
		return errors.New("can't find user")
	}

	// 判断用户房间
	if u.CurrentRoomId == 0 {
		return errors.New("user try to start game but not in room")
	}

	room, err := client.GetRoomClient().GetRoomInfo(ctx, u.CurrentRoomId)
	if err != nil {
		return err
	}

	client.GetUserCacheClient().ResetAssistNum(ctx, u.UserID)
	client.GetUserCacheClient().ResetDeadNum(ctx, u.UserID)
	client.GetUserCacheClient().ResetKillNum(ctx, u.UserID)
	client.GetUserCacheClient().SetUserIngame(ctx, u.UserID, true)
	// 给所有玩家发送他的状态
	temp := out.BuildUserReadyStatus(u.UserID, u.Currentstatus)
	for i := range room.Users {
		player := client.GetUserCacheClient().GetUserByID(ctx, room.Users[i])
		if player == nil {
			continue
		}
		rst := utils.BytesCombine(packet.BuildHeader(player.GetNextSeq(), constant.PacketTypeRoom), temp)
		packet.SendPacket(rst, player.CurrentConnection)
	}

	if room.HostUserID == u.UserID {
		// 修改房间状态
		room, err = client.GetRoomClient().StartGame(ctx, u.UserID, u.CurrentRoomId)
		if err != nil {
			return err
		}

		// 房主开始游戏
		rst := utils.BytesCombine(packet.BuildHeader(u.GetNextSeq(), constant.PacketTypeHost), out.BuildStartRoom(room.HostUserID))
		packet.SendPacket(rst, u.CurrentConnection)
		return nil
	}

	host := client.GetUserCacheClient().GetUserByID(ctx, room.HostUserID)
	if host == nil {
		client.GetUserCacheClient().SetUserIngame(ctx, u.UserID, false)
		return errors.New("can't find host")
	}

	rst := utils.BytesCombine(packet.BuildHeader(u.GetNextSeq(), constant.PacketTypeHost), out.BuildConnectHost(host.NetInfo.ExternalIpAddress, 27005))
	packet.SendPacket(rst, u.CurrentConnection)

	return nil
}
