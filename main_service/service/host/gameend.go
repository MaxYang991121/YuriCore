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

type GameEndService interface {
	Handle(ctx context.Context) error
}

type gameEndServiceImpl struct {
	client net.Conn
}

func NewGameEndService(client net.Conn) GameEndService {
	return &gameEndServiceImpl{
		client: client,
	}
}

func (g *gameEndServiceImpl) Handle(ctx context.Context) error {
	u := client.GetUserCacheClient().GetUserByConnection(ctx, g.client)
	if u == nil || u.CurrentRoomId == 0 {
		return errors.New("can't find user or user not in room")
	}

	// TODO 修改房间状态
	room, err := client.GetRoomClient().EndGame(ctx, u.UserID, u.CurrentRoomId)
	if err != nil {
		return err
	}

	// 发送结果
	setting := utils.BytesCombine([]byte{constant.OUTUpdateSettings}, out.BuildRoomSetting(room, 0xFFFFFFFF7FFFFFFF))
	hoststop := out.BuildHostStop()
	result := out.BuildRoomResult(room)
	for i := range room.Users {
		player := client.GetUserCacheClient().GetUserByID(ctx, room.Users[i])
		if player == nil {
			continue
		}
		rst := utils.BytesCombine(packet.BuildHeader(player.GetNextSeq(), constant.PacketTypeRoom), setting)
		packet.SendPacket(rst, player.CurrentConnection)

		if player.CurrentIsIngame {
			// 发送hoststop
			rst = utils.BytesCombine(packet.BuildHeader(player.GetNextSeq(), constant.PacketTypeHost), hoststop)
			packet.SendPacket(rst, player.CurrentConnection)

			// 发送结果
			rst = utils.BytesCombine(packet.BuildHeader(player.GetNextSeq(), constant.PacketTypeRoom), result)
			packet.SendPacket(rst, player.CurrentConnection)

			client.GetUserCacheClient().SetUserIngame(ctx, player.UserID, false)
		}
		client.GetUserCacheClient().SetUserStatus(ctx, player.UserID, constant.UserNotReady)

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

	return nil
}
