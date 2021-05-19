package chat

import (
	"context"
	"errors"
	"net"
	"strconv"
	"strings"

	"github.com/KouKouChan/YuriCore/main_service/client"
	"github.com/KouKouChan/YuriCore/main_service/constant"
	"github.com/KouKouChan/YuriCore/main_service/model/out"
	"github.com/KouKouChan/YuriCore/main_service/model/packet"
	"github.com/KouKouChan/YuriCore/utils"
	. "github.com/KouKouChan/YuriCore/verbose"
)

type ChatInfra interface {
	ChannelHandler(ctx context.Context) error
	RoomHandler(ctx context.Context) error
}

type chatInfraImpl struct {
	message string
	client  net.Conn
}

func GetChatInfra(data string, client net.Conn) ChatInfra {
	return &chatInfraImpl{
		message: data,
		client:  client,
	}
}

func (c *chatInfraImpl) ChannelHandler(ctx context.Context) error {
	u := client.GetUserCacheClient().GetUserByConnection(ctx, c.client)
	if u == nil {
		return errors.New("can't find user")
	}

	// 找到所有用户
	msg := out.BuildChannelMessage(u.NickName, c.message)
	users := client.GetUserCacheClient().GetChannelNoRoomUsers(ctx, u.CurrentServerIndex, u.CurrentChannelIndex)
	for i := range users {
		if users[i] == u.UserID {
			continue
		}
		player := client.GetUserCacheClient().GetUserByID(ctx, users[i])
		if player == nil {
			continue
		}
		rst := utils.BytesCombine(packet.BuildHeader(player.GetNextSeq(), constant.PacketTypeChat), msg)
		packet.SendPacket(rst, player.CurrentConnection)
	}

	rst := utils.BytesCombine(packet.BuildHeader(u.GetNextSeq(), constant.PacketTypeChat), msg)
	packet.SendPacket(rst, u.CurrentConnection)

	DebugPrintf(2, "User %+v said channel message %+v", u.UserName, c.message)
	return nil
}

func (c *chatInfraImpl) RoomHandler(ctx context.Context) error {
	u := client.GetUserCacheClient().GetUserByConnection(ctx, c.client)
	if u == nil {
		return errors.New("can't find user")
	}

	if u.CurrentRoomId == 0 {
		return errors.New("user try to speak but not in room")
	}

	room, err := client.GetRoomClient().GetRoomInfo(ctx, u.CurrentRoomId)
	if err != nil {
		return err
	}

	strs := strings.Fields(c.message)
	if len(strs) >= 2 {
		switch strs[0] {
		case "/bot":
			if room.HostUserID != u.UserID || len(strs) != 3 {
				goto nocmd
			}

			CTBot, err := strconv.Atoi(strs[1])
			if err != nil {
				msg := out.BuildRoomMessage("YuriCore", "wrong ct bot number!")
				rst := utils.BytesCombine(packet.BuildHeader(u.GetNextSeq(), constant.PacketTypeChat), msg)
				packet.SendPacket(rst, u.CurrentConnection)
				return err
			}
			strs[2] = strings.Replace(strs[2], "\x00", "", -1)
			TRBot, err := strconv.Atoi(strs[2])
			if err != nil {
				msg := out.BuildRoomMessage("YuriCore", "wrong tr bot number!")
				rst := utils.BytesCombine(packet.BuildHeader(u.GetNextSeq(), constant.PacketTypeChat), msg)
				packet.SendPacket(rst, u.CurrentConnection)
				return err
			}
			room.NumCtBots = uint8(CTBot)
			room.NumTrBots = uint8(TRBot)
		case "/money":
			if room.HostUserID != u.UserID || len(strs) != 2 {
				goto nocmd
			}

			strs[1] = strings.Replace(strs[1], "\x00", "", -1)
			Money, err := strconv.Atoi(strs[1])
			if err != nil {
				msg := out.BuildRoomMessage("YuriCore", "wrong money number!")
				rst := utils.BytesCombine(packet.BuildHeader(u.GetNextSeq(), constant.PacketTypeChat), msg)
				packet.SendPacket(rst, u.CurrentConnection)
				return err
			}
			room.StartMoney = uint16(Money)
		default:
			goto nocmd
		}

		room, err = client.GetRoomClient().UpdateRoomSafe(ctx, room)
		if err != nil {
			return err
		}
		if room == nil || room.RoomId == 0 || room.HostUserID != u.UserID {
			return errors.New("got null resp or user is not host")
		}

		// 给所有玩家发送
		setting := utils.BytesCombine([]byte{constant.OUTUpdateSettings}, out.BuildRoomSetting(room, 0xFFFFFFFF7FFFFFFF))

		for i := range room.Users {
			dest_player := client.GetUserCacheClient().GetUserByID(ctx, room.Users[i])
			if dest_player == nil {
				continue
			}
			rst := utils.BytesCombine(packet.BuildHeader(dest_player.GetNextSeq(), constant.PacketTypeRoom), setting)
			packet.SendPacket(rst, dest_player.CurrentConnection)
		}

		msg := out.BuildRoomMessage("YuriCore", "done!")
		rst := utils.BytesCombine(packet.BuildHeader(u.GetNextSeq(), constant.PacketTypeChat), msg)
		packet.SendPacket(rst, u.CurrentConnection)
		return nil
	}

nocmd:
	msg := out.BuildRoomMessage(u.NickName, c.message)
	for i := range room.Users {
		player := client.GetUserCacheClient().GetUserByID(ctx, room.Users[i])
		if player == nil {
			continue
		}
		// TODO 游戏内是否听到
		rst := utils.BytesCombine(packet.BuildHeader(player.GetNextSeq(), constant.PacketTypeChat), msg)
		packet.SendPacket(rst, player.CurrentConnection)
	}

	DebugPrintf(2, "User %+v said room message %+v", u.UserName, c.message)
	return nil
}
