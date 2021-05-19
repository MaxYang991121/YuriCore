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
	. "github.com/KouKouChan/YuriCore/verbose"
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
	if u == nil || u.CurrentRoomId == 0 {
		return errors.New("can't find user or user not in room")
	}

	// 判断用户房间
	if u.CurrentRoomId == 0 {
		return errors.New("user started game server but not in room")
	}

	room, err := client.GetRoomClient().GetRoomInfo(ctx, u.CurrentRoomId)
	if err != nil {
		return err
	}
	if room.HostUserID != u.UserID {
		return errors.New("user started game server but is not host")
	}

	// 给所有玩家发送// 给所有玩家发送
	setting := utils.BytesCombine([]byte{constant.OUTUpdateSettings}, out.BuildRoomSetting(room, 0xFFFFFFFF7FFFFFFF))
	for i := range room.Users {
		player := client.GetUserCacheClient().GetUserByID(ctx, room.Users[i])
		if player == nil {
			continue
		}

		rst := utils.BytesCombine(packet.BuildHeader(player.GetNextSeq(), constant.PacketTypeRoom), setting)
		packet.SendPacket(rst, player.CurrentConnection)

		if player.UserID != u.UserID {
			if player.IsUserReady() {
				client.GetUserCacheClient().ResetAssistNum(ctx, player.UserID)
				client.GetUserCacheClient().ResetDeadNum(ctx, player.UserID)
				client.GetUserCacheClient().ResetKillNum(ctx, player.UserID)
				client.GetUserCacheClient().SetUserIngame(ctx, player.UserID, true)
				// //给主机发送其他人的数据
				// rst := out.UDPBuild(u.GetNextSeq(), 0, player.UserID, player.NetInfo.ExternalIpAddress, 27005)
				// packet.SendPacket(rst, u.CurrentConnection)
				// //连接到主机
				// rst = out.UDPBuild(player.GetNextSeq(), 1, u.UserID, u.NetInfo.ExternalIpAddress, 27005)
				// packet.SendPacket(rst, player.CurrentConnection)
				// //加入主机
				// rst = utils.BytesCombine(packet.BuildHeader(player.GetNextSeq(), constant.PacketTypeHost), out.BuildJoinHost(u.UserID))
				// packet.SendPacket(rst, player.CurrentConnection)
				rst = utils.BytesCombine(packet.BuildHeader(player.GetNextSeq(), constant.PacketTypeHost), out.BuildConnectHost(u.NetInfo.ExternalIpAddress, 27005))
				packet.SendPacket(rst, player.CurrentConnection)
			}
		}

		temp := out.BuildUserReadyStatus(player.UserID, player.Currentstatus)
		for j := range room.Users {
			dest_player := client.GetUserCacheClient().GetUserByID(ctx, room.Users[j])
			if dest_player == nil {
				continue
			}
			rst = utils.BytesCombine(packet.BuildHeader(dest_player.GetNextSeq(), constant.PacketTypeRoom), temp)
			packet.SendPacket(rst, dest_player.CurrentConnection)
		}

	}

	DebugPrintf(2, "User %+v started game server", u.UserName)
	return nil
}
